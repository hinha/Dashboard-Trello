package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/hinha/PAM-Trello/app"
	pb "github.com/hinha/PAM-Trello/app/pb/trello"
	"github.com/hinha/PAM-Trello/app/repository/mocks"
	"github.com/hinha/PAM-Trello/app/trello"
)

func emptyTrelloData() *app.TrelloUserCard {
	return &app.TrelloUserCard{
		ID:                       "1",
		CardID:                   "1",
		CardName:                 "name",
		CardCategory:             "on_progress",
		CardVotes:                0,
		CardCheckItems:           0,
		CardCheckLists:           0,
		CardCommentCount:         0,
		CardAttachmentsCount:     0,
		CardCheckItemsComplete:   0,
		CardCheckItemsInComplete: 0,
		CardMemberName:           "-",
		CardMemberUsername:       "-",
	}
}

func TestValidTrelloShouldNil(t *testing.T) {

	valid := emptyTrelloData()

	mockTrelloRepository := new(mocks.TrelloMock)
	mockTrelloRepository.On("Store", valid).Return(valid, nil)

	service := trello.New(mockTrelloRepository)

	err := service.Create(context.TODO(), nil)
	assert.NotNil(t, err, "Error should not nil")
	assert.Equal(t, "error should be nil", err.Error())
}

func TestValidTrelloShouldAccept(t *testing.T) {
	userCardEmtpy := emptyTrelloData()

	mockTrelloRepository := new(mocks.TrelloMock)
	mockTrelloRepository.On("Store", userCardEmtpy).Return(userCardEmtpy, nil)

	service := trello.New(mockTrelloRepository)

	dummy := &pb.Response{
		Data: []*pb.Card{
			{
				CardId:               "1",
				CardCategory:         "on_progress",
				CardName:             "name",
				CardVotes:            0,
				CountCheckItems:      0,
				CountCheckLists:      0,
				CheckItemsComplete:   0,
				CheckItemsIncomplete: 0,
				CommentCount:         0,
				AttachmentsCount:     0,
				Url:                  "",
				Members:              nil,
				CreatedAt:            time.Now().Format(time.RFC3339Nano),
			},
		},
		LastUpdate: time.Now().Format(time.RFC3339Nano),
	}

	err := service.Create(context.TODO(), dummy)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, err, nil)
}
