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

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger *logrus.Entry, s Service) Service {
	return &loggingService{logger, s}
}
