package retry

import (
	"context"
	"errors"
	"github.com/demianshtepa/patterns/clock"
	clockmocks "github.com/demianshtepa/patterns/clock/mocks"
	retrymocks "github.com/demianshtepa/patterns/stability/retry/mocks"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	defaultCtx := context.Background()
	err := errors.New("fail function called")
	retryDuration := time.Second
	expectations := []struct {
		name    string
		retries int
		fn      func(context.Context) Function
		t       func() clock.Time
		ctx     func() context.Context
		result  interface{}
		err     error
	}{
		{
			name:    "without retries success result",
			retries: 0,
			fn: func(ctx context.Context) Function {
				mock := retrymocks.NewMockFunction(t)
				mock.EXPECT().Execute(ctx).Return("Ok", nil).Once()

				return mock.Execute
			},
			t: func() clock.Time {
				mock := clockmocks.NewMockTime(t)

				return mock
			},
			ctx: func() context.Context {
				return defaultCtx
			},
			result: "Ok",
			err:    nil,
		},
		{
			name:    "without retries error result",
			retries: 0,
			fn: func(ctx context.Context) Function {
				mock := retrymocks.NewMockFunction(t)
				mock.EXPECT().Execute(ctx).Return(nil, err).Once()

				return mock.Execute
			},
			t: func() clock.Time {
				mock := clockmocks.NewMockTime(t)

				return mock
			},
			ctx: func() context.Context {
				return defaultCtx
			},
			result: nil,
			err:    err,
		},
		{
			name:    "with retries success result",
			retries: 3,
			fn: func(ctx context.Context) Function {
				mock := retrymocks.NewMockFunction(t)
				mock.EXPECT().Execute(ctx).Return("Ok", nil).Once()

				return mock.Execute
			},
			t: func() clock.Time {
				mock := clockmocks.NewMockTime(t)

				return mock
			},
			ctx: func() context.Context {
				return defaultCtx
			},
			result: "Ok",
			err:    nil,
		},
		{
			name:    "with retries error result",
			retries: 3,
			fn: func(ctx context.Context) Function {
				mock := retrymocks.NewMockFunction(t)
				mock.EXPECT().Execute(ctx).Return(nil, err).Times(4)

				return mock.Execute
			},
			t: func() clock.Time {
				mockChan := make(chan time.Time, 1)
				mock := clockmocks.NewMockTime(t)
				mock.EXPECT().After(retryDuration).Return(mockChan).Run(func(_ time.Duration) {
					mockChan <- time.Now()
				}).Times(3)

				return mock
			},
			ctx: func() context.Context {
				return defaultCtx
			},
			result: nil,
			err:    err,
		},
		{
			name:    "with cancelled context",
			retries: 3,
			fn: func(ctx context.Context) Function {
				mock := retrymocks.NewMockFunction(t)
				mock.EXPECT().Execute(ctx).Return(nil, err).Once()

				return mock.Execute
			},
			t: func() clock.Time {
				mockChan := make(chan time.Time, 1)
				mock := clockmocks.NewMockTime(t)
				mock.EXPECT().After(retryDuration).Return(mockChan).Once()

				return mock
			},
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				return ctx
			},
			result: nil,
			err:    context.Canceled,
		},
	}

	for _, expectation := range expectations {
		t.Run(expectation.name, func(t *testing.T) {
			ctx := expectation.ctx()
			retry := Retry(expectation.fn(ctx), expectation.retries, expectation.t(), retryDuration)

			result, err := retry(ctx)
			if result != expectation.result {
				t.Errorf("expected result to be %s, got %s", expectation.result, result)
			}
			if err != expectation.err {
				t.Errorf("expected error to be %s, got %s", expectation.err, err)
			}
		})
	}
}
