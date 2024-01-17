package circuit_breaker

import (
	"context"
	"errors"
	"github.com/demianshtepa/patterns/clock"
	"sync"
	"time"
)

var ErrMaxAttemptsReached = errors.New("max attempts reached")

type Function func(context.Context) (interface{}, error)

func CircuitBreaker(fn Function, maxAttempts uint, t clock.Time, resetDuration time.Duration) Function {
	var attempts uint
	resetTime := t.Now()
	var mtx sync.RWMutex

	return func(ctx context.Context) (interface{}, error) {
		mtx.RLock()

		if attempts >= maxAttempts {
			if t.Now().Before(resetTime) {
				mtx.RUnlock()
				return nil, ErrMaxAttemptsReached
			}
		}
		mtx.RUnlock()

		mtx.Lock()
		defer mtx.Unlock()

		result, err := fn(ctx)
		if err != nil {
			attempts++
			resetTime = t.Now().Add(resetDuration)

			return result, err
		}
		attempts = 0

		return result, nil
	}
}
