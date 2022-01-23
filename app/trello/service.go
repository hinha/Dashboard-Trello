package trello

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/pb"
	pbTrello "github.com/hinha/PAM-Trello/app/pb/trello"
	"github.com/hinha/PAM-Trello/app/util/security"
)

type Service interface {
	Authorize(key string) (interface{}, error)
	Create(card *app.TrelloUserCard) error
	Performance(id string) (app.Performance, error)
	CardList() ([]app.TrelloUserCard, error)
	TrelloList(id string) (app.TrelloItemList, error)
	AddMember(id string, in app.TrelloAddMember) (*app.Trello, error)
	GetTotalTrello(paramYear string) (app.Performance, error)
	GetClusters(ctx context.Context, paramYear string) (app.ClusterResponse, error)
}

type service struct {
	grpcHost string
	grpcPort string

	trello  app.TrelloRepository
	account app.AccountRepository

	encrypt *security.BearerCipher
}

// Create use by cron server
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

	cards, err := s.trello.FindIDCardCategory(id)
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

	userCard, err := s.trello.FindByUserCard(id)
	perform.CardActivity = userCard

	return perform, nil
}

func (s *service) CardList() ([]app.TrelloUserCard, error) {
	return s.trello.ListCard()
}

func (s *service) TrelloList(id string) (app.TrelloItemList, error) {
	var itemList app.TrelloItemList

	account, err := s.account.ListAccount(id)
	if err != nil {
		return itemList, err
	}
	itemList.User = account

	userTrello, err := s.trello.ListTrelloUser()
	if err != nil {
		return itemList, err
	}
	itemList.TrelloUser = userTrello

	return itemList, nil
}

func (s *service) AddMember(id string, in app.TrelloAddMember) (*app.Trello, error) {
	record, err := s.trello.FindMemberID(in.MemberID)
	if err != nil {
		return nil, fmt.Errorf("error when inserted data")
	}

	if record.CardMemberID == "" {
		_, err := s.trello.StoreUser(in)

		record, err = s.trello.FindMemberID(in.MemberID)
		return record, err
	}
	return nil, fmt.Errorf("user already registered")
}

func (s *service) GetTotalTrello(paramYear string) (app.Performance, error) {
	var perform app.Performance

	category, err := s.trello.FindCardCategoryYears(paramYear)
	if err != nil {
		return perform, err
	}
	perform.CardCategory = perform.CategoryDuplicate(category)

	cards, err := s.trello.FindCategoryByYears(paramYear)
	if err != nil {
		return perform, err
	}

	lineChart := perform.LineChart(cards).JSON()
	lineChart["grid"] = map[string]interface{}{
		"left":  "3%",
		"right": "4%",
	}
	perform.Daily = lineChart

	return perform, nil
}

func (s *service) GetClusters(ctx context.Context, paramYear string) (app.ClusterResponse, error) {

	cards, err := s.trello.ListCard()
	if err != nil {
		return app.ClusterResponse{}, err
	}

	var cardAnalyze []*pbTrello.CardAnalyze
	for _, card := range cards {
		cardFormatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			card.CardCreatedAt.Year(), card.CardCreatedAt.Month(), card.CardCreatedAt.Day(),
			card.CardCreatedAt.Hour(), card.CardCreatedAt.Minute(), card.CardCreatedAt.Second())
		cardAnalyze = append(cardAnalyze, &pbTrello.CardAnalyze{
			CardId:               card.CardID,
			CardCategory:         card.CardCategory,
			CardName:             card.CardName,
			CardVotes:            card.CardVotes,
			CountCheckItems:      card.CardCheckItems,
			CountCheckLists:      card.CardCheckLists,
			CheckItemsIncomplete: card.CardCheckItemsInComplete,
			CheckItemsComplete:   card.CardCheckItemsComplete,
			CommentCount:         card.CardCommentCount,
			AttachmentsCount:     card.CardAttachmentsCount,
			Username:             card.CardMemberUsername,
			CreatedAt:            cardFormatted,
		})
	}

	client := pb.NewGrpc(s.grpcHost, s.grpcPort)
	clientService, err := pb.NewService(client.Conn).Analyze(ctx, &pbTrello.PamInput{Data: cardAnalyze})
	if err != nil {
		return app.ClusterResponse{}, err
	}

	var jsonCard interface{}
	if err := json.Unmarshal(clientService.Card, &jsonCard); err != nil {
		return app.ClusterResponse{}, err
	}

	var jsonActivity interface{}
	if err := json.Unmarshal(clientService.Activity, &jsonActivity); err != nil {
		return app.ClusterResponse{}, err
	}

	var jsonAveragePlot interface{}
	if err := json.Unmarshal(clientService.AveragePlot, &jsonAveragePlot); err != nil {
		return app.ClusterResponse{}, err
	}

	var jsonWeight interface{}
	if err := json.Unmarshal(clientService.Weight, &jsonWeight); err != nil {
		return app.ClusterResponse{}, err
	}

	response := app.ClusterResponse{
		Card:              jsonCard,
		AverageCluster:    clientService.AverageCluster,
		ScatterClustering: clientService.ScatterClustering,
		Activity:          jsonActivity,
		AveragePlot:       jsonAveragePlot,
		Weight:            jsonWeight,
	}

	return response, nil
}

func New(trello app.TrelloRepository, account app.AccountRepository) *service {
	CipherIv := "Programmer is not robot"
	CipherHeader := "sangatrahasiabro[HEHE]"
	CipherKey := "Harga kopi ditentukan oleh kualitas"
	return &service{
		grpcHost: os.Getenv("GRPC_HOST_TRELLO"),
		grpcPort: os.Getenv("GRPC_PORT_TRELLO"),

		trello:  trello,
		account: account,
		encrypt: security.NewCipher(CipherIv, CipherHeader, CipherKey),
	}
}
