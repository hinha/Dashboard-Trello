package trello

import (
	"github.com/hinha/PAM-Trello/app"
	"github.com/sirupsen/logrus"
	"time"
)

type loggingService struct {
	logger *logrus.Entry
	next   Service
}

func (s *loggingService) Create(card *app.TrelloUserCard) (err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took":    time.Since(begin),
			"data_in": card.CardID,
			"err":     err,
		}).Info("Create")
	}(time.Now())

	return s.next.Create(card)
}

func (s *loggingService) Performance(id string) (o app.Performance, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took": time.Since(begin),
			"user": id,
			"err":  err,
		}).Info("Performance")
	}(time.Now())

	return s.next.Performance(id)
}

func (s *loggingService) Authorize(key string) (o interface{}, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took":       time.Since(begin),
			"length_key": len(key),
			"err":        err,
		}).Info("Authorize")
	}(time.Now())

	return s.next.Authorize(key), err
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger *logrus.Entry, s Service) Service {
	return &loggingService{logger, s}
}
