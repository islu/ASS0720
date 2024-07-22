package user

import (
	"context"

	"github.com/islu/ASS0720/internal/domain/user"
)

// Get user tasks status by address
func (s *UserService) GetUserTaskStatus(ctx context.Context, walletAddress string) ([]user.UserTask, error) {

	// TODO: Update user status

	return s.userTaskRepo.ListUserTask_Join(ctx, walletAddress)
}

// Get user points history for distributed tasks
func (s *UserService) GetUserPointsHistory(ctx context.Context, walletAddress string) ([]user.UserTask, error) {
	return s.userTaskRepo.ListUserTask_Join(ctx, walletAddress)
}

/*
	Dashboard
*/

func (s *UserService) UpdateUniswapUSDCETHPairSwapLog(ctx context.Context, startBlockNumber, endBlockNumber int64) error {

	data, err := s.uniSwapClient.GetUniswapPairV2SwapEvent(startBlockNumber, endBlockNumber)
	if err != nil {
		return err
	}

	for _, v := range data {
		_, err = s.blockRepo.CreateUniswapUSDCETHPairSwapLog(ctx, v)
		if err != nil {
			return err
		}
	}

	return nil
}
