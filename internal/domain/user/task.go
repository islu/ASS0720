package user

import "time"

type Task struct {
	Seqno       int
	GroupNo     int
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}
