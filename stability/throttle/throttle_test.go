package throttle

import (
	"context"
	timemocks "github.com/demianshtepa/patterns/clock/mocks"
	throttlemocks "github.com/demianshtepa/patterns/stability/throttle/mocks"
	"sync"
	"testing"
	"time"
)

func TestThrottleReturnsResult(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	mockFunction := throttlemocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return("Ok", nil).Once()

	retryDuration := time.Second
	tickerChan := make(chan time.Time)
	mockTicker := timemocks.NewMockTicker(t)
	mockTicker.EXPECT().Chan().Return(tickerChan).Once()
	mockTicker.EXPECT().Stop().Run(func() {
		wg.Done()
	}).Once()
	mockTime := timemocks.NewMockTime(t)
	mockTime.EXPECT().NewTicker(retryDuration).Return(mockTicker).Once()

	throttle := Throttle(mockFunction.Execute, 5, mockTime, retryDuration)
	result, err := throttle(ctx)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if result != "Ok" {
		t.Errorf("expected result to be %s, got %s", "Ok", result)
	}

	cancel()
	wg.Wait()
}

func TestThrottleReturnsErrorMaxAttemptsReached(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	mockFunction := throttlemocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return("Ok", nil).Once()

	retryDuration := time.Second
	tickerChan := make(chan time.Time)
	mockTicker := timemocks.NewMockTicker(t)
	mockTicker.EXPECT().Chan().Return(tickerChan).Once()
	mockTicker.EXPECT().Stop().Run(func() {
		wg.Done()
	}).Once()
	mockTime := timemocks.NewMockTime(t)
	mockTime.EXPECT().NewTicker(retryDuration).Return(mockTicker).Once()

	expectations := []struct {
		result interface{}
		err    error
	}{
		{"Ok", nil},
		{nil, ErrMaxAttemptsReached},
	}

	throttle := Throttle(mockFunction.Execute, 1, mockTime, retryDuration)
	for _, expectation := range expectations {
		result, err := throttle(ctx)
		if err != expectation.err {
			t.Errorf("expected error to be %s, got %s", expectation.err, err)
		}
		if result != expectation.result {
			t.Errorf("expected result to be %s, got %s", expectation.result, result)
		}
	}

	cancel()
	wg.Wait()
}

func TestThrottleResetsAttempts(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	mockFunction := throttlemocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return("Ok", nil).Once()

	retryDuration := time.Second
	tickerChan := make(chan time.Time, 1)
	mockTicker := timemocks.NewMockTicker(t)
	mockTicker.EXPECT().Chan().Return(tickerChan)
	mockTicker.EXPECT().Stop().Run(func() {
		wg.Done()
	}).Once()
	mockTime := timemocks.NewMockTime(t)
	mockTime.EXPECT().NewTicker(retryDuration).Return(mockTicker).Once()

	throttle := Throttle(mockFunction.Execute, 1, mockTime, retryDuration)
	tickerChan <- time.Now()
	result, err := throttle(ctx)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if result != "Ok" {
		t.Errorf("expected result to be %s, got %s", "Ok", result)
	}
	tickerChan <- time.Now()
	tickerChan <- time.Now()

	cancel()
	wg.Wait()
}

func TestThrottleReturnsErrorForCancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	mockFunction := throttlemocks.NewMockFunction(t)

	retryDuration := time.Second
	mockTime := timemocks.NewMockTime(t)

	cancel()
	throttle := Throttle(mockFunction.Execute, 5, mockTime, retryDuration)
	result, err := throttle(ctx)
	if err != context.Canceled {
		t.Errorf("expected error to be %s, got %s", ErrMaxAttemptsReached, err)
	}
	if result != nil {
		t.Errorf("expected result to be nil, got %s", result)
	}

	cancel()
}
