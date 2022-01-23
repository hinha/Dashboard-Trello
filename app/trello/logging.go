package trello

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/hinha/PAM-Trello/app"
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

func (s *loggingService) TrelloList(id string) (o app.TrelloItemList, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took": time.Since(begin),
			"user": id,
			"err":  err,
		}).Info("TrelloList")
	}(time.Now())

	return s.next.TrelloList(id)
}

func (s *loggingService) CardList() (o []app.TrelloUserCard, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took":   time.Since(begin),
			"length": len(o),
			"err":    err,
		}).Info("TrelloList")
	}(time.Now())

	return s.next.CardList()
}

func (s *loggingService) AddMember(id string, in app.TrelloAddMember) (o *app.Trello, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took": time.Since(begin),
			"user": id,
			"err":  err,
		}).Info("AddMember")
	}(time.Now())

	return s.next.AddMember(id, in)
}

func (s *loggingService) Authorize(key string) (o interface{}, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took":       time.Since(begin),
			"length_key": len(key),
			"err":        err,
		}).Info("Authorize")
	}(time.Now())

	return s.next.Authorize(key)
}

func (s *loggingService) GetTotalTrello(paramYear string) (o app.Performance, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took": time.Since(begin),
			"err":  err,
		}).Info("GetClusters")
	}(time.Now())

	return s.next.GetTotalTrello(paramYear)
}

func (s *loggingService) GetClusters(ctx context.Context, paramYear string) (response app.ClusterResponse, err error) {
	defer func(begin time.Time) {
		s.logger.WithFields(logrus.Fields{
			"took": time.Since(begin),
			"err":  err,
		}).Info("GetClusters")
	}(time.Now())

	return s.next.GetClusters(ctx, paramYear)
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger *logrus.Entry, s Service) Service {
	return &loggingService{logger, s}
}
