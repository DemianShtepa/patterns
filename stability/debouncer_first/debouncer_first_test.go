package debouncer_first

import (
	"context"
	clockmocks "github.com/demianshtepa/patterns/clock/mocks"
	debouncerfirstmocks "github.com/demianshtepa/patterns/stability/debouncer_first/mocks"
	"sync"
	"testing"
	"time"
)

func TestDebounceFirstReturnsSuccessResult(t *testing.T) {
	resetDuration := time.Second
	ctx := context.Background()
	mockFunction := debouncerfirstmocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return("Ok", nil)
	mockTime := clockmocks.NewMockTime(t)
	mockTime.EXPECT().Now().Return(time.Now())
	debounce := DebounceFirst(mockFunction.Execute, mockTime, resetDuration)

	result, err := debounce(ctx)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if resultString, ok := result.(string); !ok || resultString != "Ok" {
		t.Errorf("expected string result %s, got %s", "ok", resultString)
	}
}

func TestDebounceFirstReturnsFirstResult(t *testing.T) {
	resetDuration := time.Second
	ctx := context.Background()
	mockFunction := debouncerfirstmocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return("Ok0", nil).Once()
	mockTime := clockmocks.NewMockTime(t)
	mockTime.EXPECT().Now().Return(time.Now())
	debounce := DebounceFirst(mockFunction.Execute, mockTime, resetDuration)

	expectations := [2]string{
		"Ok0",
		"Ok0",
	}

	for _, expectation := range expectations {
		result, err := debounce(ctx)
		if err != nil {
			t.Error("unexpected error", err)
		}
		if resultString, ok := result.(string); !ok || resultString != expectation {
			t.Errorf("expected string result %s, got %s", expectation, resultString)
		}
	}
}

func TestDebounceFirstResetsResult(t *testing.T) {
	resetDuration := time.Second
	ctx := context.Background()
	mockFunction := debouncerfirstmocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return("Ok0", nil).Once()
	mockFunction.EXPECT().Execute(ctx).Return("Ok1", nil).Once()
	now := time.Now()
	mockTime := clockmocks.NewMockTime(t)
	mockTime.EXPECT().Now().Return(now).Times(3)
	mockTime.EXPECT().Now().Return(now.Add(time.Hour)).Times(2)
	debounce := DebounceFirst(mockFunction.Execute, mockTime, resetDuration)

	expectations := [2]string{
		"Ok0",
		"Ok1",
	}

	for _, expectation := range expectations {
		result, err := debounce(ctx)
		if err != nil {
			t.Error("unexpected error", err)
		}
		if resultString, ok := result.(string); !ok || resultString != expectation {
			t.Errorf("expected string result %s, got %s", expectation, resultString)
		}
	}
}

func TestDebounceFirstConcurrentAccess(t *testing.T) {
	concurrentRequests := 3
	resetDuration := time.Second
	ctx := context.Background()
	mockFunction := debouncerfirstmocks.NewMockFunction(t)
	mockFunction.EXPECT().Execute(ctx).Return("Ok", nil)
	mockTime := clockmocks.NewMockTime(t)
	mockTime.EXPECT().Now().Return(time.Now())
	debounce := DebounceFirst(mockFunction.Execute, mockTime, resetDuration)
	var wg sync.WaitGroup

	wg.Add(concurrentRequests)
	for i := 0; i < concurrentRequests; i++ {
		go func() {
			defer wg.Done()

			result, err := debounce(ctx)
			if err != nil {
				t.Error("unexpected error", err)
			}
			if resultString, ok := result.(string); !ok || resultString != "Ok" {
				t.Errorf("expected string result %s, got %s", "ok", resultString)
			}
		}()
	}

	wg.Wait()
}
