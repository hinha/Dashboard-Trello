package trello

import (
	"encoding/json"
	"fmt"
	"github.com/hinha/PAM-Trello/app/util/security"

	"github.com/hinha/PAM-Trello/app"
)

type Service interface {
	Authorize(key string) (interface{}, error)
	Create(card *app.TrelloUserCard) error
	Performance(id string) (app.Performance, error)
}

type service struct {
	trello app.TrelloRepository

	encrypt *security.BearerCipher
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

func (s *service) Authorize(key string) (interface{}, error) {
	plain, err := s.encrypt.DecryptStringCBC(key)
	if err != nil {
		return nil, err
	}

	var decode map[string]interface{}
	if err := json.Unmarshal([]byte(plain), &decode); err != nil {
		return nil, err
	}

	// TODO: Need validation time expiration

	return decode, nil
}

func (s *service) Performance(id string) (app.Performance, error) {
	var perform app.Performance

	doneTask, err := s.trello.FindCardCategory(id, "DONE")
	if err != nil {
		return perform, err
	}
	perform.Done = doneTask

	progressTask, err := s.trello.FindCardCategory(id, "ON PROGRESS")
	if err != nil {
		return perform, nil
	}
	perform.OnProgress = progressTask

	todoTask, err := s.trello.FindCardCategory(id, "TODO")
	if err != nil {
		return perform, nil
	}
	perform.Todo = todoTask

	// TODO: create chart visual daily

	return perform, nil
}

func New(trello app.TrelloRepository) *service {
	CipherIv := "Programmer is not robot"
	CipherHeader := "sangatrahasiabro[HEHE]"
	CipherKey := "Harga kopi ditentukan oleh kualitas"
	return &service{trello: trello, encrypt: security.NewCipher(CipherIv, CipherHeader, CipherKey)}
}
