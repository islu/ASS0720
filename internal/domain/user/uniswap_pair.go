package user

type UniswapPairSwapEvent struct {
	From            string // Sender address
	BlockNumber     uint64
	TransactionHash string
	Timestamp       uint64
	Amount0In       uint64
	Amount0Out      uint64
	Amount1Out      uint64
	Amount1In       uint64
}
