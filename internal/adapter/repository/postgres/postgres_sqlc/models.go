// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package postgres_sqlc

import (
	"time"
)

type Task struct {
	Seqno       int32
	TaskGroupNo int32
	TaskName    string
	TaskDesc    string
	StartTime   time.Time
	EndTime     time.Time
}

type UniswapUsdcEthPairSwapLog struct {
	BlockSender string
	BlockNumber int64
	BlockTime   time.Time
	TxHash      string
	Amount0In   int64
	Amount0Out  int64
	Amount1In   int64
	Amount1Out  int64
}

type UserTask struct {
	TaskSeqno     int32
	WalletAddress string
	TotalAmount   int64
	Point         int32
	Status        string
	CreateTime    time.Time
	UpdateTime    time.Time
}
