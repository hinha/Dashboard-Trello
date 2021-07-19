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

func (s *loggingService) AuthLogin(ctx context.Context, in *app.LoginInput) (o *app.Accounts, token string, err error) {
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

func (s *loggingService) NewAccount(ctx context.Context, adminID string, roleName string, in *app.RegisterInput) (err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"method":        "register",
			"took":          time.Since(begin),
			"admin_id":      adminID,
			"authorize":     roleName,
			"register_name": in.Name,
			"err":           err,
		}).Info("NewAccount")
	}(time.Now())

	return s.next.NewAccount(ctx, adminID, roleName, in)
}

func (s *loggingService) ListAccount(ctx context.Context, adminID string, roleName string) (o []app.Accounts, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"method":     "register",
			"took":       time.Since(begin),
			"admin_id":   adminID,
			"authorize":  roleName,
			"count_data": len(o),
			"err":        err,
		}).Info("NewAccount")
	}(time.Now())

	return s.next.ListAccount(ctx, adminID, roleName)
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger *logrus.Entry, s Service) Service {
	return &loggingService{logger, s}
}
