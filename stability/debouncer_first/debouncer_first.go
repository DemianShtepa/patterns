package debouncer_first

import (
	"context"
	"sync"
	"time"
)

type TimeProvider func() time.Time

type Function func(context.Context) (interface{}, error)

func DebounceFirst(fn Function, now TimeProvider, resetDuration time.Duration) Function {
	var result interface{}
	var err error
	var mx sync.Mutex
	resetTime := now()
	return func(ctx context.Context) (interface{}, error) {
		mx.Lock()
		defer mx.Unlock()

		if now().Before(resetTime) {
			return result, err
		}

		result, err = fn(ctx)
		resetTime = now().Add(resetDuration)

		return result, err
	}
}
