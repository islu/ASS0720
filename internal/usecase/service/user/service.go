package user

import "context"

type UserService struct {
	userTaskRepo  UserTaskRepository
	blockRepo     BlockRepository
	uniSwapClient UniswapClient
}

type UserServiceParam struct {
	UserTaskRepo  UserTaskRepository
	BlockRepo     BlockRepository
	UniswapClient UniswapClient
}

func NewUserService(_ context.Context, param UserServiceParam) *UserService {
	return &UserService{
		userTaskRepo:  param.UserTaskRepo,
		blockRepo:     param.BlockRepo,
		uniSwapClient: param.UniswapClient,
	}
}
