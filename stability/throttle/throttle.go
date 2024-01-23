package throttle

import (
	"context"
	"errors"
	"github.com/demianshtepa/patterns/clock"
	"sync"
	"time"
)

var ErrMaxAttemptsReached = errors.New("max attempts reached")

type Function func(context.Context) (interface{}, error)

func Throttle(fn Function, maxAttempts int, t clock.Time, retryDuration time.Duration) Function {
	var attempts = maxAttempts
	var once sync.Once
	var mtx sync.Mutex

	return func(ctx context.Context) (interface{}, error) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		once.Do(func() {
			ticker := t.NewTicker(retryDuration)

			go func() {
				defer ticker.Stop()

				for {
					select {
					case <-ticker.Chan():
						mtx.Lock()
						attempts++
						if attempts > maxAttempts {
							attempts = maxAttempts
						}
						mtx.Unlock()
					case <-ctx.Done():
						return
					}
				}
			}()
		})

		mtx.Lock()
		defer mtx.Unlock()

		if attempts == 0 {
			return nil, ErrMaxAttemptsReached
		}

		attempts--

		return fn(ctx)
	}
}
