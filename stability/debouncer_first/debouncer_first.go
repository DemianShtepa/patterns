package debouncer_first

import (
	"context"
	"github.com/demianshtepa/patterns/clock"
	"sync"
	"time"
)

type Function func(context.Context) (interface{}, error)

func DebounceFirst(fn Function, t clock.Time, resetDuration time.Duration) Function {
	var result interface{}
	var err error
	var mx sync.Mutex
	resetTime := t.Now()
	return func(ctx context.Context) (interface{}, error) {
		mx.Lock()
		defer mx.Unlock()

		if t.Now().Before(resetTime) {
			return result, err
		}

		result, err = fn(ctx)
		resetTime = t.Now().Add(resetDuration)

		return result, err
	}
}
