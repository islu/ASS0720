package user

import (
	"context"

	"github.com/islu/ASS0720/internal/domain/user"
)

type UserTaskRepository interface {
	// Get user points history for distributed tasks
	ListUserTask_Join(ctx context.Context, walletAddress string) ([]user.UserTask, error)
}

type UniswapClient interface {
	// Debug print
	DebugPrint_UniswapPairV2SwapEvent()
}
