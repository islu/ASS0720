package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/islu/HW0720/internal/domain/user"
)

const (
	// Contract address for Uniswap V2 Pair
	uniswapPairV2Address = "0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"

	// ABI for Uniswap V2 Pair (simplified)
	uniswapV2PairABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount0In","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1In","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount0Out","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1Out","type":"uint256"},{"indexed":true,"internalType":"address","name":"to","type":"address"}],"name":"Swap","type":"event"}]`

	// Swap event signature
	swapEventSig = "Swap(address,uint256,uint256,uint256,uint256,address)"
)

func (c *EthereumClient) GetUniswapPairV2SwapEvent(fromBlockNumber, toBlockNumber int64) ([]user.UniswapPairSwapEvent, error) {

	// Connect to json rpc node
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/" + c.AlchemyAPIKey)
	if err != nil {
		return nil, err
	}

	// Contract address
	contractAddress := common.HexToAddress(uniswapPairV2Address)

	// Swap event signature
	swapEventSignature := []byte(swapEventSig)
	swapEventHash := crypto.Keccak256Hash(swapEventSignature)

	fromBlock := big.NewInt(fromBlockNumber)
	toBlock := big.NewInt(toBlockNumber)

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{swapEventHash},
		},
	}

	// Get logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV2PairABI))
	if err != nil {
		return nil, err
	}

	var swapEvents []user.UniswapPairSwapEvent
	for i := 0; i < len(logs); i++ {
		vLog := logs[i]

		blockNumber := big.NewInt(int64(vLog.BlockNumber))
		txIndex := vLog.TxIndex

		// Get block
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatalf("Failed to get block: %v", err)
		}
		// Get transaction
		tx := block.Transactions()[txIndex]

		from, err := types.Sender(types.NewLondonSigner(tx.ChainId()), tx)
		if err != nil {
			return nil, err
		}

		// Decode event data
		event := struct {
			Sender     common.Address
			Amount0In  *big.Int
			Amount1In  *big.Int
			Amount0Out *big.Int
			Amount1Out *big.Int
			To         common.Address
		}{}

		err = parsedABI.UnpackIntoInterface(&event, "Swap", vLog.Data)
		if err != nil {
			return nil, err
		}

		// Decode Topics
		event.Sender = common.HexToAddress(vLog.Topics[1].Hex())
		event.To = common.HexToAddress(vLog.Topics[2].Hex())

		swapEvents = append(swapEvents, user.UniswapPairSwapEvent{
			From:            from.Hex(),
			BlockNumber:     vLog.BlockNumber,
			TransactionHash: vLog.TxHash.Hex(),
			Timestamp:       block.Time(),
			Amount0In:       event.Amount0In.Uint64(),
			Amount0Out:      event.Amount0Out.Uint64(),
			Amount1Out:      event.Amount1Out.Uint64(),
			Amount1In:       event.Amount1In.Uint64(),
		})
	}

	return swapEvents, nil
}

func (c *EthereumClient) DebugPrint_UniswapPairV2SwapEvent() {
	// Connect to json rpc node
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/" + c.AlchemyAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	// Contract address
	contractAddress := common.HexToAddress(uniswapPairV2Address)

	// Swap event signature
	swapEventSignature := []byte(swapEventSig)
	swapEventHash := crypto.Keccak256Hash(swapEventSignature)

	// fromTimestamp := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC).Unix()
	// fromBlock, err := getBlockNumberByTimestamp(client, fromTimestamp)
	// if err != nil {
	// 	log.Fatalf("Failed to get block number by timestamp: %v", err)
	// }

	// toTimestamp := time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC).Unix()
	// toBlock, err := getBlockNumberByTimestamp(client, toTimestamp)
	// if err != nil {
	// 	log.Fatalf("Failed to get block number by timestamp: %v", err)
	// }

	fromBlock := big.NewInt(20352000)
	// toBlock := big.NewInt(20353000)

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		// ToBlock:   toBlock,
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{swapEventHash},
		},
	}

	// Get logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV2PairABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Print logs
	logCount := len(logs)
	start := 0
	if logCount > 5 {
		start = logCount - 5
	}
	fmt.Println(logCount)
	for i := start; i < logCount; i++ {
		vLog := logs[i]
		// fmt.Printf("Log: %+v\n", vLog)

		fmt.Printf("Address: %s\n", vLog.Address.Hex())
		fmt.Printf("Block Hash: %s\n", vLog.BlockHash.Hex())
		fmt.Printf("Index: %d\n", vLog.Index)
		fmt.Printf("Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Transaction Index: %d\n", vLog.TxIndex)
		fmt.Printf("Transaction Hash: %s\n", vLog.TxHash.Hex())
		fmt.Println()

		blockNumber := big.NewInt(int64(vLog.BlockNumber))
		txIndex := vLog.TxIndex

		// Get block
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatalf("Failed to get block: %v", err)
		}
		// Get transaction
		tx := block.Transactions()[txIndex]

		fmt.Printf("Timestamp: %d\n", block.Time())
		if from, err := types.Sender(types.NewLondonSigner(tx.ChainId()), tx); err == nil {
			fmt.Printf("From(Sender): %s\n", from.Hex())
		}

		fmt.Println()

		// Decode event data
		event := struct {
			Sender     common.Address
			Amount0In  *big.Int
			Amount1In  *big.Int
			Amount0Out *big.Int
			Amount1Out *big.Int
			To         common.Address
		}{}

		err = parsedABI.UnpackIntoInterface(&event, "Swap", vLog.Data)
		if err != nil {
			log.Fatalf("Failed to unpack log data: %v", err)
		}

		// Decode Topics
		event.Sender = common.HexToAddress(vLog.Topics[1].Hex())
		event.To = common.HexToAddress(vLog.Topics[2].Hex())

		// Print event data
		fmt.Printf("Sender: %s\n", event.Sender.Hex())
		fmt.Printf("Amount0In: %s\n", event.Amount0In.String()) // USDC
		fmt.Printf("Amount0Out: %s\n", event.Amount0Out.String())
		fmt.Printf("Amount1In: %s\n", event.Amount1In.String()) // WETH
		fmt.Printf("Amount1Out: %s\n", event.Amount1Out.String())
		fmt.Printf("To: %s\n", event.To.Hex())

		fmt.Println()
		fmt.Println("---------------------------")
		fmt.Println()
	}
}

func getBlockNumberByTimestamp(client *ethclient.Client, timestamp int64) (*big.Int, error) {
	// Binary search
	var startBlock, endBlock uint64 = 0, 99999999
	var startTime, endTime int64

	// Get start block timestamp
	startBlockHeader, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(startBlock)))
	if err != nil {
		return nil, err
	}
	startTime = int64(startBlockHeader.Time)

	// Get end block timestamp
	endBlockHeader, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(endBlock)))
	if err != nil {
		return nil, err
	}
	endTime = int64(endBlockHeader.Time)

	for startBlock <= endBlock {
		midBlock := (startBlock + endBlock) / 2
		midBlockHeader, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(midBlock)))
		if err != nil {
			return nil, err
		}
		midTime := int64(midBlockHeader.Time)

		if midTime < timestamp {
			startBlock = midBlock + 1
			startTime = midTime
		} else if midTime > timestamp {
			endBlock = midBlock - 1
			endTime = midTime
		} else {
			return big.NewInt(int64(midBlock)), nil
		}
	}

	// Return the closest block
	if timestamp-startTime < endTime-timestamp {
		return big.NewInt(int64(startBlock)), nil
	}
	return big.NewInt(int64(endBlock)), nil
}
