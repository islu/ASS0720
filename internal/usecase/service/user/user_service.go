package user

import (
	"context"

	"github.com/islu/ASS0720/internal/domain/user"
)

// Get user points history for distributed tasks
func (s *UserService) GetUserTaskList(ctx context.Context) ([]user.UserTask, error) {
	return s.userTaskRepo.GetUserTaskList(ctx)
}
