package postgres

import (
	"context"
	"time"

	psqlc "github.com/islu/HW0720/internal/adapter/repository/postgres/postgres_sqlc"
	"github.com/islu/HW0720/internal/domain/user"
)

// Create Uniswap USDC/ETH pair swap log
func (r *PostgresRepository) CreateUniswapUSDCETHPairSwapLog(ctx context.Context, params user.UniswapPairSwapEvent) (*user.UniswapPairSwapEvent, error) {

	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	q := psqlc.New(r.connPool)
	qtx := q.WithTx(tx)

	log, err := qtx.CreateUniswapUSDCETHPairSwapLog(ctx, psqlc.CreateUniswapUSDCETHPairSwapLogParams{
		BlockSender: params.From,
		BlockNumber: int64(params.BlockNumber),
		BlockTime:   params.Timestamp,
		TxHash:      params.TransactionHash,
		Amount0In:   int64(params.Amount0In),
		Amount0Out:  int64(params.Amount0Out),
		Amount1In:   int64(params.Amount1In),
		Amount1Out:  int64(params.Amount1Out),
	})

	if err != nil {
		return nil, err
	}

	result := buildUniswapPairSwapEvent(log)

	return &result, tx.Commit(ctx)
}

// List Uniswap USDC/ETH pair swap log by sender
func (r *PostgresRepository) ListUniswapUSDCETHPairSwapLogBySender(ctx context.Context, sender string) ([]user.UniswapPairSwapEvent, error) {

	q := psqlc.New(r.connPool)

	list, err := q.ListUniswapUSDCETHPairSwapLogBySender(ctx, sender)
	if err != nil {
		return nil, err
	}

	var result []user.UniswapPairSwapEvent
	for _, log := range list {
		result = append(result, buildUniswapPairSwapEvent(log))
	}

	return result, nil
}

// List Uniswap USDC/ETH pair swap log between timestamp
func (r *PostgresRepository) ListUniswapUSDCETHPairSwapLogByTimestamp(ctx context.Context, startTime, endTime time.Time) ([]user.UniswapPairSwapEvent, error) {

	q := psqlc.New(r.connPool)

	list, err := q.ListUniswapUSDCETHPairSwapLogByTimestamp(ctx, psqlc.ListUniswapUSDCETHPairSwapLogByTimestampParams{
		BlockTime:   startTime,
		BlockTime_2: endTime,
	})
	if err != nil {
		return nil, err
	}

	var result []user.UniswapPairSwapEvent
	for _, log := range list {
		result = append(result, buildUniswapPairSwapEvent(log))
	}

	return result, nil
}

func buildUniswapPairSwapEvent(log psqlc.UniswapUsdcEthPairSwapLog) user.UniswapPairSwapEvent {

	return user.UniswapPairSwapEvent{
		From:            log.BlockSender,
		BlockNumber:     uint64(log.BlockNumber),
		TransactionHash: log.TxHash,
		Timestamp:       log.BlockTime,
		Amount0In:       uint64(log.Amount0In),
		Amount0Out:      uint64(log.Amount0Out),
		Amount1In:       uint64(log.Amount1In),
		Amount1Out:      uint64(log.Amount1Out),
	}
}
