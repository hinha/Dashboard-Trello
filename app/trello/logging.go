package trello

import (
	"context"
	"github.com/hinha/PAM-Trello/app/pb/trello"
	"github.com/sirupsen/logrus"
	"time"
)

type loggingService struct {
	logger *logrus.Entry
	next   Service
}

func (s *loggingService) Create(ctx context.Context, mod *trello.Response) (err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"method":     "accessControl",
			"took":       time.Since(begin),
			"data_in":    mod.LastUpdate,
			"data_error": mod.Error,
			"err":        err,
		}).Info("DeleteAccount")
	}(time.Now())

	return s.next.Create(ctx, mod)
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger *logrus.Entry, s Service) Service {
	return &loggingService{logger, s}
}
