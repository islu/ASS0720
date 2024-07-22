package blockchain

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetUniswapPairV2SwapEvent(t *testing.T) {

	client := createEthereumClient(t)

	data, err := client.GetUniswapPairV2SwapEvent(20358617, 20358638)
	assert.NoError(t, err)
	assert.Equal(t, 11, len(data))

	// 20358617 // https://etherscan.io/tx/0xddadd3a76ff526dedcb6d7d182f77bc7aeeae46900d988e3e2734e5b4b7ffa0e
	assert.Equal(t, uint64(20358617), data[0].BlockNumber)
	assert.Equal(t, "0x9057Cb12392539C553ebE2148627F3D79f310553", data[0].From)
	assert.Equal(t, "0xddadd3a76ff526dedcb6d7d182f77bc7aeeae46900d988e3e2734e5b4b7ffa0e", data[0].TransactionHash)
	assert.Equal(t, uint64(1721609627), data[0].Timestamp)
	assert.Equal(t, uint64(0), data[0].Amount0In)
	assert.Equal(t, uint64(137186002), data[0].Amount0Out)
	assert.Equal(t, uint64(38753682456344840), data[0].Amount1In)
	assert.Equal(t, uint64(0), data[0].Amount1Out)

	// 20358638 // https://etherscan.io/tx/0xf79d5b9d1364a712ab4fd1768fb1b9311ba1c3d3afd58ad1035c26180ee1cd97
	assert.Equal(t, uint64(20358638), data[10].BlockNumber)
	assert.Equal(t, "0xFAdB20fF8d6ff919ccf57C1C22aacCFfA422292f", data[10].From)
	assert.Equal(t, "0xf79d5b9d1364a712ab4fd1768fb1b9311ba1c3d3afd58ad1035c26180ee1cd97", data[10].TransactionHash)
	assert.Equal(t, uint64(1721609879), data[10].Timestamp)
	assert.Equal(t, uint64(100000000), data[10].Amount0In)
	assert.Equal(t, uint64(0), data[10].Amount0Out)
	assert.Equal(t, uint64(0), data[10].Amount1In)
	assert.Equal(t, uint64(28075764597248419), data[10].Amount1Out)
}

func createEthereumClient(t *testing.T) *EthereumClient {

	err := godotenv.Load("../../../.env")
	assert.NoError(t, err)

	apiKey := os.Getenv("ALCHEMY_API_KEY")
	assert.NotEmpty(t, apiKey)

	client := &EthereumClient{
		Env:           "test",
		AlchemyAPIKey: apiKey,
	}
	return client
}
