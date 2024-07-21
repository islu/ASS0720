package user

import (
	"context"

	"github.com/islu/ASS0720/internal/domain/user"
)

// Get user points history for distributed tasks
func (s *UserService) GetUserPointsHistory(ctx context.Context, walletAddress string) ([]user.UserTask, error) {
	return s.userTaskRepo.ListUserTask_Join(ctx, walletAddress)
}
