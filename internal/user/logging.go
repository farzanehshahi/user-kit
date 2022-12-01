package user

import (
	"context"
	"github.com/farzanehshahi/user-kit/internal/entity"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	next   Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (ls *loggingService) Create(ctx context.Context, reqUser *entity.User) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.next.Create(ctx, reqUser)
}

func (ls *loggingService) Get(ctx context.Context, id string) (user entity.User, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "get",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.next.Get(ctx, id)
}

func (ls *loggingService) Update(ctx context.Context, id, username, password string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "update",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.next.Update(ctx, id, username, password)
}

func (ls *loggingService) Delete(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			"method", "delete",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return ls.next.Delete(ctx, id)
}
