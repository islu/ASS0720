package user

import "time"

type UniswapPairSwapEvent struct {
	From            string // Sender address
	BlockNumber     uint64
	TransactionHash string
	Timestamp       time.Time
	Amount0In       uint64
	Amount0Out      uint64
	Amount1Out      uint64
	Amount1In       uint64
}
