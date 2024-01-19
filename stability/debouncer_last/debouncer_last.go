package debouncer_last

import (
	"context"
	"github.com/demianshtepa/patterns/clock"
	"sync"
	"time"
)

type Function func(context.Context) (interface{}, error)

func DebounceLast(fn Function, t clock.Time, thresholdDuration time.Duration) Function {
	var threshold = t.Now()
	var result interface{}
	var err error
	var once sync.Once
	var m sync.Mutex

	return func(ctx context.Context) (interface{}, error) {
		m.Lock()
		defer m.Unlock()

		threshold = t.Now().Add(thresholdDuration)

		once.Do(func() {
			ticker := t.NewTicker(time.Millisecond * 100)
			go func() {
				defer func() {
					m.Lock()
					ticker.Stop()
					once = sync.Once{}
					m.Unlock()
				}()
				for {
					select {
					case <-ticker.Chan():
						m.Lock()
						if t.Now().After(threshold) {
							result, err = fn(ctx)
							m.Unlock()
							return
						}
						m.Unlock()
					case <-ctx.Done():
						m.Lock()
						result, err = "", ctx.Err()
						m.Unlock()
						return
					}
				}
			}()
		})

		return result, err
	}
}
