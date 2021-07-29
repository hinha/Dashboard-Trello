package trello

import (
	"fmt"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/util/security"
)

type Service interface {
	Authorize(key string) (interface{}, error)
	Create(card *app.TrelloUserCard) error
	Performance(id string) (app.Performance, error)
}

type service struct {
	trello  app.TrelloRepository
	account app.AccountRepository

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
	return security.Authorize(s.encrypt, key)
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

	usrOnline, _ := s.account.GetOnlineStatus(id)
	var onlineUsers []map[string]interface{}
	for _, user := range usrOnline {
		onlineUsers = append(onlineUsers, map[string]interface{}{
			"last_active": user.LastLogin,
			"username":    user.Username,
			"name":        user.Name,
		})
	}
	perform.OnlineUsers = onlineUsers

	return perform, nil
}

func New(trello app.TrelloRepository, account app.AccountRepository) *service {
	CipherIv := "Programmer is not robot"
	CipherHeader := "sangatrahasiabro[HEHE]"
	CipherKey := "Harga kopi ditentukan oleh kualitas"
	return &service{trello: trello, account: account, encrypt: security.NewCipher(CipherIv, CipherHeader, CipherKey)}
}
