package user

import (
	"context"
	"github.com/farzanehshahi/user-kit/internal/entity"
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram

	//	countResult metrics.Histogram

	next Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}

func (is instrumentingService) Create(ctx context.Context, reqUser *entity.User) (err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", "create"}
		is.requestCount.With(lvs...).Add(1)
		is.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.next.Create(ctx, reqUser)
}

func (is instrumentingService) Get(ctx context.Context, id string) (user entity.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get"}
		is.requestCount.With(lvs...).Add(1)
		is.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.next.Get(ctx, id)
}

func (is instrumentingService) Update(ctx context.Context, id, username, password string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "update"}
		is.requestCount.With(lvs...).Add(1)
		is.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.next.Update(ctx, id, username, password)
}

func (is instrumentingService) Delete(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "delete"}
		is.requestCount.With(lvs...).Add(1)
		is.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return is.next.Delete(ctx, id)
}
