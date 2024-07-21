package user

import "context"

type UserService struct {
	userTaskRepo  UserTaskRepository
	uniSwapClient UniswapClient
}

type UserServiceParam struct {
	UserTaskRepo  UserTaskRepository
	UniswapClient UniswapClient
}

func NewUserService(_ context.Context, param UserServiceParam) *UserService {
	return &UserService{
		userTaskRepo:  param.UserTaskRepo,
		uniSwapClient: param.UniswapClient,
	}
}
