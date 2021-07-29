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

	dup := func(sample []app.CardCategory) []app.CardCategory {
		var unique []app.CardCategory

	sampleLoop:
		for _, v := range sample {
			for i, u := range unique {
				if v.Label == u.Label {
					unique[i] = v
					continue sampleLoop
				}
			}
			unique = append(unique, v)
		}
		return unique
	}
	var perform app.Performance

	cards, err := s.trello.FindCardCategory(id)
	if err != nil {
		return perform, err
	}

	perform.CardCategory = dup(cards)

	groupBy, err := s.trello.CategoryByDate(id)
	if err != nil {
		return perform, err
	}

	lineChart := perform.LineChart(groupBy).JSON()
	lineChart["grid"] = map[string]interface{}{
		"left":  "3%",
		"right": "4%",
	}
	perform.Daily = lineChart

	pieChart := perform.PieChart(groupBy).JSON()
	pieChart["grid"] = map[string]interface{}{
		"left":  "3%",
		"right": "4%",
	}
	perform.Task = pieChart
	// TODO: create chart visual daily

	return perform, nil
}

func New(trello app.TrelloRepository) *service {
	CipherIv := "Programmer is not robot"
	CipherHeader := "sangatrahasiabro[HEHE]"
	CipherKey := "Harga kopi ditentukan oleh kualitas"
	return &service{trello: trello, encrypt: security.NewCipher(CipherIv, CipherHeader, CipherKey)}
}
