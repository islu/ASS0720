package user

import (
	"context"
	"time"

	"github.com/islu/HW0720/internal/domain/user"
)

type UserTaskRepository interface {
	// Create task
	CreateTask(ctx context.Context, params user.Task) (*user.Task, error)
	// List task
	ListTask(ctx context.Context) ([]user.Task, error)
	// List task by group number
	ListTaskByGroupNo(ctx context.Context, groupNo int) ([]user.Task, error)

	// Create user task
	CreateUserTask(ctx context.Context, params user.UserTask) (*user.UserTask, error)
	// List user task
	ListUserTask_Join(ctx context.Context, walletAddress string) ([]user.UserTask, error)
}

type BlockRepository interface {
	CreateUniswapUSDCETHPairSwapLog(ctx context.Context, params user.UniswapPairSwapEvent) (*user.UniswapPairSwapEvent, error)
	ListUniswapUSDCETHPairSwapLogBySender(ctx context.Context, sender string) ([]user.UniswapPairSwapEvent, error)
	ListUniswapUSDCETHPairSwapLogByTimestamp(ctx context.Context, startTime, endTime time.Time) ([]user.UniswapPairSwapEvent, error)
}

type UniswapClient interface {
	GetUniswapPairV2SwapEvent(fromBlockNumber, toBlockNumber int64) ([]user.UniswapPairSwapEvent, error)

	DebugPrint_UniswapPairV2SwapEvent() // Debug print
}
