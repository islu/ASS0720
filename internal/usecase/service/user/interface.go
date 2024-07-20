package user

import (
	"context"

	"github.com/islu/ASS0720/internal/domain/user"
)

type UserTaskRepository interface {
	// Get user points history for distributed tasks
	GetUserTaskList(ctx context.Context) ([]user.UserTask, error)
}
