package clock

import "time"

type Ticker interface {
	Chan() <-chan time.Time
	Reset(time.Duration)
	Stop()
}

type realTicker struct {
	*time.Ticker
}

func NewTicker(d time.Duration) Ticker {
	return &realTicker{time.NewTicker(d)}
}

func (r *realTicker) Chan() <-chan time.Time {
	return r.C
}
