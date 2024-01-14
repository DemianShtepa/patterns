package circuit_breaker

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrMaxAttemptsReached = errors.New("max attempts reached")

type TimeProvider func() time.Time

type Function func(context.Context) (interface{}, error)

func CircuitBreaker(fn Function, maxAttempts uint, now TimeProvider, resetDuration time.Duration) Function {
	var attempts uint
	resetTime := now()
	var mtx sync.RWMutex

	return func(ctx context.Context) (interface{}, error) {
		mtx.RLock()

		if attempts >= maxAttempts {
			if now().Before(resetTime) {
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
			resetTime = now().Add(resetDuration)

			return result, err
		}
		attempts = 0

		return result, nil
	}
}
