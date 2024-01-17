package circuit_breaker

import (
	"context"
	"errors"
	clockmocks "github.com/demianshtepa/patterns/clock/mocks"
	circuitbreakermocks "github.com/demianshtepa/patterns/stability/circuit_breaker/mocks"
	"sync"
	"testing"
	"time"
)

func TestCircuitBreakerDoesntProduceErrors(t *testing.T) {
	maxAttempts := 3
	resetDuration := time.Second
	ctx := context.Background()
	mockNow := clockmocks.NewMockTime(t)
	mockNow.EXPECT().Now().Return(time.Now())
	mockFunction := circuitbreakermocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return(nil, nil)
	circuitBreaker := CircuitBreaker(mockFunction.Execute, uint(maxAttempts), mockNow, resetDuration)

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
	ctx := context.Background()
	mockNow := clockmocks.NewMockTime(t)
	mockNow.EXPECT().Now().Return(time.Now())
	mockFunction := circuitbreakermocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return(nil, errors.New("fail function called"))
	circuitBreaker := CircuitBreaker(mockFunction.Execute, uint(maxAttempts), mockNow, resetDuration)
	var err error

	for i := 0; i <= maxAttempts; i++ {
		_, err = circuitBreaker(ctx)
		if err == nil {
			t.Error("expected error", err)
		}
	}

	if !errors.Is(err, ErrMaxAttemptsReached) {
		t.Error("expected error", ErrMaxAttemptsReached)
	}
}

func TestCircuitBreakerResetsAttempts(t *testing.T) {
	maxAttempts := 3
	resetDuration := time.Second
	now := time.Now()
	ctx := context.Background()
	mockNow := clockmocks.NewMockTime(t)
	mockNow.EXPECT().Now().Return(now).Times(4)
	mockNow.EXPECT().Now().Return(now.Add(time.Hour)).Times(2)
	mockFunction := circuitbreakermocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return(nil, errors.New("fail function called"))
	circuitBreaker := CircuitBreaker(mockFunction.Execute, uint(maxAttempts), mockNow, resetDuration)
	var err error

	for i := 0; i < maxAttempts; i++ {
		_, err = circuitBreaker(ctx)
		if err == nil {
			t.Error("expected error", err)
		}
	}

	_, err = circuitBreaker(ctx)
	if errors.Is(err, ErrMaxAttemptsReached) {
		t.Error("unexpected error", err)
	}
}

func TestCircuitBreakerConcurrentAccess(t *testing.T) {
	maxAttempts := 3
	resetDuration := time.Second
	ctx := context.Background()
	mockNow := clockmocks.NewMockTime(t)
	mockNow.EXPECT().Now().Return(time.Now())
	mockFunction := circuitbreakermocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return(nil, nil)
	circuitBreaker := CircuitBreaker(mockFunction.Execute, uint(maxAttempts), mockNow, resetDuration)
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
