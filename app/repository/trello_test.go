package repository_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/repository/mocks"
	"github.com/hinha/PAM-Trello/app/trello"
)

func emptyTrelloData() *app.TrelloUserCard {
	return &app.TrelloUserCard{
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
		CardMemberName:           "",
		CardMemberUsername:       "",
	}
}

func TestValidTrelloShouldNil(t *testing.T) {

	valid := emptyTrelloData()

	mockTrelloRepository := new(mocks.TrelloMock)
	mockTrelloRepository.On("Store", valid).Return(valid, nil)

	service := trello.New(mockTrelloRepository)

	err := service.Create(nil)
	assert.NotNil(t, err, "Error should not nil")
	assert.Equal(t, "error should be nil", err.Error())
}

func TestValidTrelloShouldAccept(t *testing.T) {
	userCardEmtpy := emptyTrelloData()

	mockTrelloRepository := new(mocks.TrelloMock)
	mockTrelloRepository.On("Store", userCardEmtpy).Return(userCardEmtpy, nil)

	service := trello.New(mockTrelloRepository)

	err := service.Create(userCardEmtpy)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, err, nil)
}
