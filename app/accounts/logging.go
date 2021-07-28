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

func (s *loggingService) GetProfile(ctx context.Context, id string) (o *app.Accounts, secret string, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"method":     "account",
			"took":       time.Since(begin),
			"request_id": id,
			"username":   o.Username,
			"err":        err,
		}).Info("GetProfile")
	}(time.Now())

	return s.next.GetProfile(ctx, id)
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
			"method":     "list",
			"took":       time.Since(begin),
			"admin_id":   adminID,
			"authorize":  roleName,
			"count_data": len(o),
			"err":        err,
		}).Info("ListAccount")
	}(time.Now())

	return s.next.ListAccount(ctx, adminID, roleName)
}

func (s *loggingService) DeleteAccount(ctx context.Context, adminId string, roleName string, userID string, userName string) (err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"method":     "delete",
			"took":       time.Since(begin),
			"admin_id":   adminId,
			"authorize":  roleName,
			"tag_delete": userName,
			"err":        err,
		}).Info("DeleteAccount")
	}(time.Now())

	return s.next.DeleteAccount(ctx, adminId, roleName, userID, userName)
}

func (s *loggingService) GetAccessList(ctx context.Context) (o app.AccessControl, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"method": "control",
			"took":   time.Since(begin),
			"err":    err,
		}).Info("GetAccessList")
	}(time.Now())

	return s.next.GetAccessList(ctx)
}

func (s *loggingService) NewAccessControlList(ctx context.Context, adminId string, roleAdmin string, control *app.AssignRole) (err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"method":    "accessControl",
			"took":      time.Since(begin),
			"admin_id":  adminId,
			"authorize": roleAdmin,
			"err":       err,
		}).Info("DeleteAccount")
	}(time.Now())

	return s.next.NewAccessControlList(ctx, adminId, roleAdmin, control)
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger *logrus.Entry, s Service) Service {
	return &loggingService{logger, s}
}
