package user

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/islu/HW0720/internal/domain/user"
)

// Get user tasks status by address
func (s *UserService) GetUserTaskStatus(ctx context.Context, walletAddress string) ([]user.UserTask, error) {

	// Get user tasks by address
	tasks, err := s.userTaskRepo.ListUserTask_Join(ctx, walletAddress)
	if err != nil {
		err = errors.Join(errors.New("[UserService][GetUserTaskStatus] Get user tasks failed"), err)
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("[UserService][GetUserTaskStatus] User has no tasks")
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].TaskSeqno < tasks[j].TaskSeqno
	})
	// Get onboarding task
	onboardingTask := tasks[0]
	// Get share pool task // Assumption: Rest of tasks are share pool tasks
	sharePoolTask := tasks[1:]

	// Get swap logs from blockchain (DB cache)
	currTime := time.Now()
	logs, err := s.blockRepo.ListUniswapUSDCETHPairSwapLogByTimestamp(ctx, onboardingTask.TaskStartTime, currTime)
	if err != nil {
		err = errors.Join(errors.New("[UserService][GetUserTaskStatus] Get swap logs failed"), err)
		return nil, err
	}

	// Update onboarding task status
	var updatedTasks []user.UserTask

	updatedOnboardingTask := user.UpdateOnboardingTaskStatus(onboardingTask, currTime, logs)
	updatedTasks = append(updatedTasks, updatedOnboardingTask)

	// Update share pool task status
	for _, share := range sharePoolTask {
		updatedSharePoolTask := user.UpdateSharePoolTaskStatus(share, onboardingTask, currTime, logs)
		updatedTasks = append(updatedTasks, updatedSharePoolTask)
	}

	// Store updated tasks
	var storeTasks []user.UserTask
	for _, task := range updatedTasks {
		updated, err := s.userTaskRepo.UpdateUserTask(ctx, walletAddress, task)
		if err != nil {
			err = errors.Join(errors.New("[UserService][GetUserTaskStatus] Update user tasks failed"), err)
			return nil, err
		}
		storeTasks = append(storeTasks, *updated)
	}

	return storeTasks, nil
}

// Get user points history for distributed tasks
func (s *UserService) GetUserPointsHistory(ctx context.Context, walletAddress string) ([]user.UserTask, error) {

	tasks, err := s.userTaskRepo.ListUserTask_Join(ctx, walletAddress)
	if err != nil {
		return nil, err
	}

	// Filter tasks that are completed
	var completedTasks []user.UserTask
	for _, task := range tasks {
		if task.Status == user.UserTaskStatus_Claimed {
			completedTasks = append(completedTasks, task)
		}
	}

	return completedTasks, nil
}

// Distribute tasks for user
func (s *UserService) DistributeTasks(ctx context.Context, walletAddress string) error {

	// Get user tasks by address
	tasks, err := s.userTaskRepo.ListUserTask_Join(ctx, walletAddress)

	if err != nil {
		return err
	}

	// If user has no tasks, initialize them
	if len(tasks) == 0 {

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
				Status:        user.UserTaskStatus_NotStarted,
			})
			if err != nil {
				return err
			}
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
		err = errors.Join(errors.New("[UserService][UpdateUniswapUSDCETHPairSwapLog] Get swap logs failed"), err)
		return err
	}

	for _, v := range data {
		_, err = s.blockRepo.CreateUniswapUSDCETHPairSwapLog(ctx, v)
		if err != nil {
			err = errors.Join(errors.New("[UserService][UpdateUniswapUSDCETHPairSwapLog] Create swap logs failed"), err)
			return err
		}
	}

	return nil
}
