package trello

import (
	"fmt"

	"github.com/hinha/PAM-Trello/app"
)

type Service interface {
	Create(card *app.TrelloUserCard) error
}

type service struct {
	trello app.TrelloRepository
}

func (s *service) Create(card *app.TrelloUserCard) error {
	if card == nil {
		return fmt.Errorf("error should be nil")
	}

	_, err := s.trello.Store(card)
	if err != nil {
		return err
	}

	return nil
}

func New(trello app.TrelloRepository) *service {
	return &service{trello: trello}
}
