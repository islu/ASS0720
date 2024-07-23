package user

import "time"

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

var (
	UserTaskStatus_NoStarted  = "no_started"
	UserTaskStatus_InProgress = "in_progress"
	UserTaskStatus_Outdated   = "outdated"
	UserTaskStatus_Claimed    = "claimed"
)

// Check if onboarding task is completed by wallet address.
// Onboarding task is completed if user swap at least 1000u
//
// Assumption:
//   - Every buy and sell is tracked
//   - Accumulated over time, not based on single trades
func CheckIsCompleteOnboardingTask(onboardingTask UserTask, source []UniswapPairSwapEvent) bool {
	var swapU uint64
	for _, s := range source {
		if onboardingTask.TaskStartTime.Before(s.Timestamp) && s.From == onboardingTask.WalletAddress {
			swapU = swapU + s.Amount0In + s.Amount0Out
			if swapU >= 1000 {
				return true
			}
		}
	}
	return false
}

func GetSharePoolTaskEarnPoint(sharePoolTask UserTask, source []UniswapPairSwapEvent) int {
	var totalU, userU uint64

	for _, s := range source {
		if sharePoolTask.TaskStartTime.Before(s.Timestamp) && s.From == sharePoolTask.WalletAddress {
			userU = userU + s.Amount0In + s.Amount0Out
		}
		totalU = totalU + s.Amount0In + s.Amount0Out

		if sharePoolTask.TaskEndTime.Before(s.Timestamp) {
			break
		}
	}

	point := float64(userU) / float64(totalU) * float64(1000)

	return int(point)
}
