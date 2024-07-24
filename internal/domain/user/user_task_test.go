package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateOnboardingTaskStatus(t *testing.T) {

	testCases := []struct {
		source              []UniswapPairSwapEvent
		target              UserTask
		expectedStatus      string
		expectedPoints      int
		expectedTotalAmount int64
	}{
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(1 * time.Hour),
					Amount0In:  990000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
				Points:        0,
				TotalAmount:   0,
				Status:        UserTaskStatus_NotStarted,
			},
			UserTaskStatus_InProgress,
			0,
			int64(990),
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(1 * time.Hour),
					Amount0In:  990000000,
					Amount0Out: 0,
				},
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(2 * time.Hour),
					Amount0In:  990000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
				Points:        0,
				TotalAmount:   0,
				Status:        UserTaskStatus_NotStarted,
			},
			UserTaskStatus_Claimed,
			100,
			int64(1980),
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(1 * time.Hour),
					Amount0In:  990000000,
					Amount0Out: 0,
				},
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(2 * time.Hour),
					Amount0In:  0,
					Amount0Out: 990000000,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
				Points:        0,
				TotalAmount:   0,
				Status:        UserTaskStatus_InProgress,
			},
			UserTaskStatus_Claimed,
			100,
			int64(1980),
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(1 * time.Hour),
					Amount0In:  990000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now().Add(-48 * time.Hour),
				TaskEndTime:   time.Now().Add(-24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
				Points:        0,
				TotalAmount:   0,
				Status:        UserTaskStatus_NotStarted,
			},
			UserTaskStatus_Outdated,
			0,
			int64(0),
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(-10 * time.Hour),
					Amount0In:  990000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now().Add(-48 * time.Hour),
				TaskEndTime:   time.Now().Add(-1 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
				Points:        0,
				TotalAmount:   0,
				Status:        UserTaskStatus_InProgress,
			},
			UserTaskStatus_Claimed,
			0,
			int64(990),
		},
	}

	for _, e := range testCases {
		updated := UpdateOnboardingTaskStatus(e.target, time.Now(), e.source)
		assert.Equal(t, e.expectedStatus, updated.Status)
		assert.Equal(t, e.expectedPoints, updated.Points)
		assert.Equal(t, e.expectedTotalAmount, updated.TotalAmount)
	}
}

func TestUpdateSharePoolTaskStatus(t *testing.T) {
	testCases := []struct {
		source              []UniswapPairSwapEvent
		target              UserTask
		onboardingTask      UserTask
		expectedStatus      string
		expectedTotalAmount int64
		expectedPoints      int
	}{
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(2 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
				{
					From:       "0x566375059f718fAA640823c5D5Dac34ED5B1ac56",
					Timestamp:  time.Now().Add(4 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
				{
					From:       "0x566375059f718fAA640823c5D5Dac34ED5B1ac56",
					Timestamp:  time.Now().Add(6 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now().Add(-48 * time.Hour),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
				Points:        0,
				TotalAmount:   0,
				Status:        UserTaskStatus_NotStarted,
			},
			UserTask{
				Status: UserTaskStatus_Claimed,
				Points: 100,
			},
			UserTaskStatus_InProgress,
			int64(1000),
			3333,
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(-2 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
				{
					From:       "0x566375059f718fAA640823c5D5Dac34ED5B1ac56",
					Timestamp:  time.Now().Add(-4 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
				{
					From:       "0x566375059f718fAA640823c5D5Dac34ED5B1ac56",
					Timestamp:  time.Now().Add(-6 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now().Add(-48 * time.Hour),
				TaskEndTime:   time.Now().Add(-1 * time.Minute),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
				Points:        0,
				TotalAmount:   0,
				Status:        UserTaskStatus_InProgress,
			},
			UserTask{
				Status: UserTaskStatus_Claimed,
				Points: 100,
			},
			UserTaskStatus_Claimed,
			int64(1000),
			3333,
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(-2 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
				{
					From:       "0x566375059f718fAA640823c5D5Dac34ED5B1ac56",
					Timestamp:  time.Now().Add(-4 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
				{
					From:       "0x566375059f718fAA640823c5D5Dac34ED5B1ac56",
					Timestamp:  time.Now().Add(-6 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now().Add(-48 * time.Hour),
				TaskEndTime:   time.Now().Add(-1 * time.Minute),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
				Points:        0,
				TotalAmount:   0,
				Status:        UserTaskStatus_InProgress,
			},
			UserTask{
				Status: UserTaskStatus_InProgress,
				Points: 0,
			},
			UserTaskStatus_Claimed,
			int64(1000),
			0,
		},
	}

	for _, e := range testCases {
		updated := UpdateSharePoolTaskStatus(e.target, e.onboardingTask, time.Now(), e.source)
		assert.Equal(t, e.expectedStatus, updated.Status)
		assert.Equal(t, e.expectedPoints, updated.Points)
		assert.Equal(t, e.expectedTotalAmount, updated.TotalAmount)
	}
}

func TestCheckIsCompleteOnboardingTask(t *testing.T) {

	testCases := []struct {
		source           []UniswapPairSwapEvent
		target           UserTask
		expectedComplete bool
		expectedSwapU    int64
	}{
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(15 * time.Hour),
					Amount0In:  990000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
			},
			false,
			int64(990),
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(15 * time.Hour),
					Amount0In:  9990000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
			},
			true,
			int64(9990),
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(15 * time.Hour),
					Amount0In:  0,
					Amount0Out: 9990000000,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
			},
			true,
			int64(9990),
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(15 * time.Hour),
					Amount0In:  0,
					Amount0Out: 990000000,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
			},
			false,
			int64(990),
		},
	}

	for _, e := range testCases {
		complete, swapU := CheckIsCompleteOnboardingTask(e.target, e.source)
		assert.Equal(t, e.expectedComplete, complete)
		assert.Equal(t, e.expectedSwapU, swapU)
	}
}

func TestGetSharePoolTaskEarnPoint(t *testing.T) {

	testCases := []struct {
		source         []UniswapPairSwapEvent
		target         UserTask
		expectedSwapU  int64
		expectedPoints int
	}{
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(15 * time.Hour),
					Amount0In:  990000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
			},
			int64(990),
			10000,
		},
		{
			[]UniswapPairSwapEvent{
				{
					From:       "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
					Timestamp:  time.Now().Add(2 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
				{
					From:       "0x566375059f718fAA640823c5D5Dac34ED5B1ac56",
					Timestamp:  time.Now().Add(4 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
				{
					From:       "0x566375059f718fAA640823c5D5Dac34ED5B1ac56",
					Timestamp:  time.Now().Add(6 * time.Minute),
					Amount0In:  1000_000000,
					Amount0Out: 0,
				},
			},
			UserTask{
				TaskStartTime: time.Now(),
				TaskEndTime:   time.Now().Add(24 * time.Hour),
				WalletAddress: "0xb36ce3F521a1FF4B0CffB82DDd71F0a51fCf56ce",
			},
			int64(1000),
			3333,
		},
	}

	for _, e := range testCases {
		swapU, points := GetSharePoolTaskEarnPoint(e.target, e.source)
		assert.Equal(t, e.expectedSwapU, swapU)
		assert.Equal(t, e.expectedPoints, points)
	}
}
