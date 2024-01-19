package retry

import (
	"context"
	"github.com/demianshtepa/patterns/clock"
	"time"
)

type Function func(context.Context) (interface{}, error)

func Retry(fn Function, retries int, t clock.Time, retryDuration time.Duration) Function {
	return func(ctx context.Context) (interface{}, error) {
		for i := 0; ; i++ {
			result, err := fn(ctx)
			if err == nil || i >= retries {
				return result, err
			}

			select {
			case <-t.After(retryDuration):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
	}
}
