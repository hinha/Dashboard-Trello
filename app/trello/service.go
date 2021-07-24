package trello

import (
	"context"
	"fmt"
	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/pb/trello"
	"strings"
	"time"
)

type Service interface {
	Create(ctx context.Context, mod *trello.Response) error
}

type service struct {
	trello app.TrelloRepository
}

func (s *service) Create(ctx context.Context, mod *trello.Response) error {
	if mod == nil {
		return fmt.Errorf("error should be nil")
	}

	for _, card := range mod.Data {

		var memberName []string
		//var memberID []string
		var memberUsername []string
		for _, member := range card.Members {
			memberName = append(memberName, member.Name)
			//memberID = append(memberID, member.Id)
			memberUsername = append(memberUsername, member.Username)
			//memberName = strings.Join(, ":")
		}

		span, _ := time.Parse("2013-04-01 22:43:22", card.CreatedAt)

		memberNameStr := strings.Join(memberName, ":")
		if memberNameStr == "" {
			memberNameStr = "-"
		}

		memberUserStr := strings.Join(memberUsername, ":")
		if memberUserStr == "" {
			memberUserStr = "-"
		}

		model := &app.TrelloUserCard{
			ID:                       card.CardId,
			CardID:                   card.CardId,
			CardName:                 card.CardName,
			CardCategory:             card.CardCategory,
			CardVotes:                card.CardVotes,
			CardCheckItems:           card.CountCheckItems,
			CardCheckLists:           card.CountCheckLists,
			CardCommentCount:         card.CommentCount,
			CardAttachmentsCount:     card.AttachmentsCount,
			CardCheckItemsComplete:   card.CheckItemsComplete,
			CardCheckItemsInComplete: card.CheckItemsIncomplete,
			CardMemberName:           memberNameStr,
			CardMemberUsername:       memberUserStr,
			CardCreatedAt:            span,
		}
		_, err := s.trello.Store(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func New(trello app.TrelloRepository) *service {
	return &service{trello: trello}
}
