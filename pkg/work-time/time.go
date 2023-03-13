package work_time

import "time"

func NowPointer() *time.Time {
	now := time.Now()
	return &now
}
