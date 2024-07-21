package user

import "time"

type UserTask struct {
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
