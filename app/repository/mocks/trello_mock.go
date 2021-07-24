package mocks

import (
	"github.com/hinha/PAM-Trello/app"
	"github.com/stretchr/testify/mock"
)

type TrelloMock struct {
	mock.Mock
}

func (_m *TrelloMock) Store(input *app.TrelloUserCard) (*app.TrelloUserCard, error) {
	ret := _m.Called(input)

	var r0 *app.TrelloUserCard
	if rf, ok := ret.Get(0).(func(card *app.TrelloUserCard) *app.TrelloUserCard); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(*app.TrelloUserCard)
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(card *app.TrelloUserCard) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
