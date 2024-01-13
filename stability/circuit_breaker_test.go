package stability_test

import (
	"context"
	"errors"
	"patterns/stability"
	"sync"
	"testing"
	"time"
)

var (
	successFunction stability.Function = func(ctx context.Context) (interface{}, error) {
		return nil, nil
	}
	failFunction stability.Function = func(ctx context.Context) (interface{}, error) {
		return nil, errors.New("fail function called")
	}
	currentTimeProvider stability.TimeProvider = func() time.Time {
		return time.Now()
	}
	mockTimeProvider = func(maxAttempts uint, mainTime, fallbackTime time.Time) stability.TimeProvider {
		var attempt uint = 0
		return func() time.Time {
			attempt++

			if attempt >= maxAttempts+2 {
				return fallbackTime
			}

			return mainTime
		}
	}
)

func TestCircuitBreakerDoesntProduceErrors(t *testing.T) {
	maxAttempts := 3
	resetDuration := time.Second
	circuitBreaker := stability.CircuitBreaker(successFunction, uint(maxAttempts), currentTimeProvider, resetDuration)
	ctx := context.Background()

	for i := 0; i <= maxAttempts; i++ {
		_, err := circuitBreaker(ctx)
		if err != nil {
			t.Error("unexpected error", err)
		}
	}
}

func TestCircuitBreakerProducesErrors(t *testing.T) {
	maxAttempts := 3
	resetDuration := time.Second
	circuitBreaker := stability.CircuitBreaker(failFunction, uint(maxAttempts), currentTimeProvider, resetDuration)
	ctx := context.Background()
	var err error

	for i := 0; i <= maxAttempts; i++ {
		_, err = circuitBreaker(ctx)
		if err == nil {
			t.Error("expected error", err)
		}
	}

	if !errors.Is(err, stability.ErrMaxAttemptsReached) {
		t.Error("expected error", stability.ErrMaxAttemptsReached)
	}
}

func TestCircuitBreakerResetsAttempts(t *testing.T) {
	maxAttempts := 3
	resetDuration := time.Second
	mockTime := mockTimeProvider(uint(maxAttempts), time.Now().Add(time.Hour*(-1)), time.Now())
	circuitBreaker := stability.CircuitBreaker(failFunction, uint(maxAttempts), mockTime, resetDuration)
	ctx := context.Background()
	var err error

	for i := 0; i < maxAttempts; i++ {
		_, err = circuitBreaker(ctx)
		if err == nil {
			t.Error("expected error", err)
		}
	}

	_, err = circuitBreaker(ctx)
	if errors.Is(err, stability.ErrMaxAttemptsReached) {
		t.Error("unexpected error", err)
	}
}

func TestConcurrentCircuitBreakerAccess(t *testing.T) {
	maxAttempts := 3
	resetDuration := time.Second
	circuitBreaker := stability.CircuitBreaker(successFunction, uint(maxAttempts), currentTimeProvider, resetDuration)
	ctx := context.Background()
	var wg sync.WaitGroup

	wg.Add(maxAttempts)
	for i := 0; i < maxAttempts; i++ {
		go func() {
			defer wg.Done()

			_, err := circuitBreaker(ctx)
			if err != nil {
				t.Error("unexpected error", err)
			}
		}()
	}

	wg.Wait()
}
