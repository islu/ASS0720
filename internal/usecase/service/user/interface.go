package user

import (
	"context"
	"time"

	"github.com/islu/HW0720/internal/domain/user"
)

type UserTaskRepository interface {
	// Get user points history for distributed tasks
	ListUserTask_Join(ctx context.Context, walletAddress string) ([]user.UserTask, error)
}

type BlockRepository interface {
	CreateUniswapUSDCETHPairSwapLog(ctx context.Context, params user.UniswapPairSwapEvent) (*user.UniswapPairSwapEvent, error)
	ListUniswapUSDCETHPairSwapLogBySender(ctx context.Context, sender string) ([]user.UniswapPairSwapEvent, error)
	ListUniswapUSDCETHPairSwapLogByTimestamp(ctx context.Context, startTime, endTime time.Time) ([]user.UniswapPairSwapEvent, error)
}

type UniswapClient interface {
	GetUniswapPairV2SwapEvent(fromBlockNumber, toBlockNumber int64) ([]user.UniswapPairSwapEvent, error)
	// Debug print
	DebugPrint_UniswapPairV2SwapEvent()
}
