package user

import "time"

type Task struct {
	GroupNo     int
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}
