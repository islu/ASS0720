package user

import (
	"fmt"
	"time"
)

type UserTask struct {
	TaskSeqno       int
	TaskName        string
	TaskDescription string
	TaskStartTime   time.Time
	TaskEndTime     time.Time
	WalletAddress   string
	Points          int
	TotalAmount     int64
	Status          string
	CreateTime      time.Time
	UpdateTime      time.Time
}

// Status flow:
//   - not_started ---> not_started
//   - not_started ---> in_progress
//   - not_started ---> outdated
//   - in_progress ---> in_progress
//   - in_progress ---> claimed
//   - outdated ---> outdated
//   - claimed ---> claimed

var (
	UserTaskStatus_NotStarted = "not_started"
	UserTaskStatus_InProgress = "in_progress"
	UserTaskStatus_Outdated   = "outdated"
	UserTaskStatus_Claimed    = "claimed"
)

// Update onboarding task status
func UpdateOnboardingTaskStatus(onboardingTask UserTask, currTime time.Time, source []UniswapPairSwapEvent) UserTask {

	if onboardingTask.Status == UserTaskStatus_Claimed {
		return onboardingTask
	}
	if onboardingTask.Status == UserTaskStatus_Outdated {
		return onboardingTask
	}

	updated := onboardingTask
	updated.UpdateTime = currTime

	if onboardingTask.TaskStartTime.After(currTime) {
		updated.Status = UserTaskStatus_NotStarted
		return updated
	}
	if onboardingTask.TaskEndTime.Before(currTime) {
		updated.Status = UserTaskStatus_Outdated
		return updated
	}
	updated.Status = UserTaskStatus_InProgress

	if updated.Status == UserTaskStatus_InProgress {
		if ok, swapU := CheckIsCompleteOnboardingTask(onboardingTask, source); ok {
			updated.Status = UserTaskStatus_Claimed
			updated.Points = 100
			updated.TotalAmount = swapU
		} else {
			updated.Status = UserTaskStatus_InProgress
			updated.Points = 0
			updated.TotalAmount = swapU
		}

		if onboardingTask.TaskEndTime.Before(currTime) {
			updated.Status = UserTaskStatus_Claimed
		}
	}

	return updated
}

// Update share pool task status
func UpdateSharePoolTaskStatus(onboardingTask, sharePoolTask UserTask, currTime time.Time, source []UniswapPairSwapEvent) UserTask {

	if sharePoolTask.Status == UserTaskStatus_Claimed {
		return sharePoolTask
	}
	if sharePoolTask.Status == UserTaskStatus_Outdated {
		return sharePoolTask
	}

	updated := sharePoolTask
	updated.UpdateTime = currTime

	if sharePoolTask.TaskStartTime.After(currTime) {
		updated.Status = UserTaskStatus_NotStarted
		return updated
	}
	if sharePoolTask.TaskEndTime.Before(currTime) {
		updated.Status = UserTaskStatus_Outdated
		return updated
	}
	updated.Status = UserTaskStatus_InProgress

	if updated.Status == UserTaskStatus_InProgress {

		usdc, points := GetSharePoolTaskEarnPoint(sharePoolTask, source)

		fmt.Println(updated.TaskName, usdc, points)

		// Check if onboarding task is completed
		if onboardingTask.Status == UserTaskStatus_Claimed {
			updated.Points = points
			updated.TotalAmount = usdc
		} else {
			updated.Points = 0
			updated.TotalAmount = usdc
		}

		if sharePoolTask.TaskEndTime.Before(currTime) {
			updated.Status = UserTaskStatus_Claimed
		}
	}

	return updated
}

// Check if onboarding task is completed by wallet address.
// Onboarding task is completed if user swap at least 1000u
//
// Assumption:
//   - Every buy and sell is tracked
//   - Accumulated over time, not based on single trades
func CheckIsCompleteOnboardingTask(onboardingTask UserTask, source []UniswapPairSwapEvent) (bool, int64) {
	var swapU uint64
	for _, s := range source {
		if onboardingTask.TaskStartTime.Before(s.Timestamp) &&
			onboardingTask.TaskEndTime.After(s.Timestamp) &&
			s.From == onboardingTask.WalletAddress {

			swapU = swapU + s.Amount0In + s.Amount0Out
		}
	}

	usdc := int64(swapU) / 1000000
	if swapU >= 1000_000000 {
		return true, usdc
	}
	return false, usdc
}

// Get share pool task earn point
//
// Return:
//   - User swap volume (USDC)
//   - User share pool points
func GetSharePoolTaskEarnPoint(sharePoolTask UserTask, source []UniswapPairSwapEvent) (int64, int) {
	var totalU, userU uint64

	for _, s := range source {
		if sharePoolTask.TaskStartTime.Before(s.Timestamp) && sharePoolTask.TaskEndTime.After(s.Timestamp) {

			if s.From == sharePoolTask.WalletAddress {
				userU = userU + s.Amount0In + s.Amount0Out
			}

			totalU = totalU + s.Amount0In + s.Amount0Out
		}
	}

	usdc := int64(userU) / 1000000

	point := float64(userU) / float64(totalU) * float64(10000)

	return usdc, int(point)
}
