package clock

import "time"

type Time interface {
	Now() time.Time
	NewTicker(time.Duration) Ticker
	Sleep(time.Duration)
	After(time.Duration) <-chan time.Time
}
