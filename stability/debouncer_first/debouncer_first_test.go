package debouncer_first

import (
	"context"
	"sync"
	"testing"
	"time"
)

var (
	function Function = func(ctx context.Context) (interface{}, error) {
		return "Ok", nil
	}
	mockFunction = func() Function {
		calls := [2]string{
			"Ok0",
			"Ok1",
		}
		var call int

		return func(ctx context.Context) (interface{}, error) {
			result := calls[call]
			call++

			return result, nil
		}
	}
	currentTimeProvider TimeProvider = func() time.Time {
		return time.Now()
	}
	mockTimeProvider = func() TimeProvider {
		var attempt uint = 0
		mainTime := time.Now()
		fallbackTime := mainTime.Add(time.Hour)
		return func() time.Time {
			attempt++

			if attempt >= 4 {
				return fallbackTime
			}

			return mainTime
		}
	}
)

func TestDebounceFirstReturnsSuccessResult(t *testing.T) {
	resetDuration := time.Second
	debounce := DebounceFirst(function, currentTimeProvider, resetDuration)
	ctx := context.Background()

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

	debounce := DebounceFirst(mockFunction(), currentTimeProvider, resetDuration)
	ctx := context.Background()

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

	debounce := DebounceFirst(mockFunction(), mockTimeProvider(), resetDuration)
	ctx := context.Background()

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
	debounce := DebounceFirst(function, currentTimeProvider, resetDuration)
	ctx := context.Background()
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
