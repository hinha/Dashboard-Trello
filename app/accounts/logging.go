package accounts

import (
	"context"
	"github.com/hinha/PAM-Trello/app"
	"github.com/sirupsen/logrus"
	"time"
)

type loggingService struct {
	logger *logrus.Entry
	next   Service
}

func (s *loggingService) AuthLogin(ctx context.Context, in *app.LoginInput) (token string, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"method":   "authenticate",
			"took":     time.Since(begin),
			"username": in.Username,
			"err":      err,
		}).Info("AuthLogin")
	}(time.Now())

	return s.next.AuthLogin(ctx, in)
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger *logrus.Entry, s Service) Service {
	return &loggingService{logger, s}
}
