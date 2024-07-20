package user

import "context"

type UserService struct {
	userTaskRe UserTaskRepository
}

type UserServiceParam struct {
	UserTaskRepo UserTaskRepository
}

func NewUserService(_ context.Context, param UserServiceParam) *UserService {
	return &UserService{
		userTaskRe: param.UserTaskRepo,
	}
}
