package clock

import "time"

type Time interface {
	Now() time.Time
}

type realTime struct{}

func (rt realTime) Now() time.Time {
	return time.Now()
}

func NewTime() Time {
	return realTime{}
}
