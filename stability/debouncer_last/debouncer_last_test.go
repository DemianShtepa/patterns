package debouncer_last

import (
	"context"
	clockmocks "github.com/demianshtepa/patterns/clock/mocks"
	debouncerlastmocks "github.com/demianshtepa/patterns/stability/debouncer_last/mocks"
	"sync"
	"testing"
	"time"
)

func TestDebounceLastReturnsEmptyFirstResult(t *testing.T) {
	thresholdDuration := time.Second
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	mockFunction := debouncerlastmocks.NewMockFunction(t)
	mockTicker := clockmocks.NewMockTicker(t)
	mockTicker.EXPECT().Chan().Return(make(<-chan time.Time)).Once()
	mockTicker.EXPECT().Stop().Run(func() {
		wg.Done()
	})
	mockTime := clockmocks.NewMockTime(t)
	mockTime.EXPECT().Now().Return(time.Now())
	mockTime.EXPECT().NewTicker(time.Millisecond * 100).Return(mockTicker)
	debounce := DebounceLast(mockFunction.Execute, mockTime, thresholdDuration)

	result, err := debounce(ctx)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if result != nil {
		t.Error("expected result to be nil", result)
	}
	cancel()
	wg.Wait()
}

func TestDebounceLastReturnsSecondCallResult(t *testing.T) {
	thresholdDuration := time.Second
	ctx := context.Background()
	var wg sync.WaitGroup
	mockFunction := debouncerlastmocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return("Ok", nil).Times(2)
	now := time.Now()
	tickerChan := make(chan time.Time)
	mockTicker := clockmocks.NewMockTicker(t)
	mockTicker.EXPECT().Chan().Return(tickerChan)
	mockTicker.EXPECT().Stop().Run(func() {
		wg.Done()
	})
	mockTime := clockmocks.NewMockTime(t)
	mockTime.EXPECT().NewTicker(time.Millisecond * 100).Return(mockTicker)
	mockTime.EXPECT().Now().Return(now).Times(3)
	mockTime.EXPECT().Now().Return(now.Add(time.Hour)).Once()
	mockTime.EXPECT().Now().Return(now).Times(1)
	mockTime.EXPECT().Now().Return(now.Add(time.Hour)).Once()
	debounce := DebounceLast(mockFunction.Execute, mockTime, thresholdDuration)

	expectations := []struct {
		result interface{}
		err    error
	}{
		{
			result: nil,
			err:    nil,
		},
		{
			result: "Ok",
			err:    nil,
		},
	}

	for _, expectation := range expectations {
		wg.Add(1)
		result, err := debounce(ctx)
		if result != expectation.result {
			t.Errorf("expected result to be %s, got %s", expectation.result, result)
		}
		if err != expectation.err {
			t.Errorf("expected error to be %s, got %s", expectation.err, err)
		}

		tickerChan <- time.Now()
		if expectation.result == nil {
			tickerChan <- time.Now()
		}

		wg.Wait()
	}
}
