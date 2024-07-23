package user

import (
	"context"
	"time"

	"github.com/islu/HW0720/internal/domain/user"
)

// Get user tasks status by address
func (s *UserService) GetUserTaskStatus(ctx context.Context, walletAddress string) ([]user.UserTask, error) {

	// Get user tasks by address
	tasks, err := s.userTaskRepo.ListUserTask_Join(ctx, walletAddress)
	if err != nil {
		return nil, err
	}

	// If user has no tasks, initialize them
	if len(tasks) == 0 {
		err = s.DistributeTasks(ctx, walletAddress)
		if err != nil {
			return nil, err
		}
	}

	// Update user tasks status
	// err = s.userTaskRepo.UpdateUserTaskStatus(ctx, walletAddress)
	// if err != nil {
	// 	return nil, err
	// }

	return tasks, nil
}

// Get user points history for distributed tasks
func (s *UserService) GetUserPointsHistory(ctx context.Context, walletAddress string) ([]user.UserTask, error) {
	return s.userTaskRepo.ListUserTask_Join(ctx, walletAddress)
}

// Distribute tasks for user
func (s *UserService) DistributeTasks(ctx context.Context, walletAddress string) error {

	// NOTE: Use all tasks simply for now
	tasks, err := s.userTaskRepo.ListTask(ctx)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		_, err = s.userTaskRepo.CreateUserTask(ctx, user.UserTask{
			WalletAddress: walletAddress,
			TaskSeqno:     task.Seqno,
			Points:        0,
			TotalAmount:   0,
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
			Status:        user.UserTaskStatus_NoStarted,
		})
		if err != nil {
			return err
		}
	}

	return nil
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
