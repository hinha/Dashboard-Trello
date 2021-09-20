package cron_server

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocraft/work"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/hinha/PAM-Trello/app"
	pb "github.com/hinha/PAM-Trello/app/pb/trello"
	"github.com/hinha/PAM-Trello/app/trello"
)

type trelloJobHandler struct {
	s trello.Service

	logger *log.Entry
}

func (h *trelloJobHandler) name() string {
	return "trello_stream"
}

func (h *trelloJobHandler) middleware(ctx *work.Job, next work.NextMiddlewareFunc) error {
	return next()
}

func (h *trelloJobHandler) job(c *work.Job) error {

	provider, err := h.client()
	if err != nil {
		h.logger.Error(err)
		return err
	}

	response, err := provider.GetCard(context.Background(), &pb.Request{
		ApiKey: os.Getenv("TRELLO_KEY"),
		Token:  os.Getenv("TRELLO_TOKEN"),
		Board:  &pb.Board{Id: "5d4bf62b1ee5a58abc7aae1b", Name: "Kalkula"},
	})
	if err != nil {
		h.logger.Error(err)
		return err
	}

	if response.Error != "" {
		h.logger.Error(response.Error)
		return fmt.Errorf(response.Error)
	}

	for _, card := range response.Data {

		span, err := strconv.ParseInt(strconv.FormatInt(card.CreatedAt, 10), 10, 64)
		if err != nil {
			h.logger.Error(err)
			return err
		}

		model := &app.TrelloUserCard{
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
			CardMemberID:             card.MemberId,
			CardMemberName:           card.MemberName,
			CardMemberUsername:       card.MemberUsername,
			CardUrl:                  card.Url,
			CardCreatedAt:            time.Unix(span, 0),
		}
		if err := h.s.Create(model); err != nil {
			h.logger.Error(err)
		}
	}
	return nil
}

func (h *trelloJobHandler) client() (pb.TrelloClient, error) {

	portLine := fmt.Sprintf("%s:%s", os.Getenv("GRPC_HOST_TRELLO"), os.Getenv("GRPC_PORT_TRELLO"))
	client, err := grpc.Dial(portLine,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*10)),
		grpc.WithInsecure())

	if err != nil {
		return nil, err
	}
	h.logger.Infof("connected grpc on %s", portLine)

	return pb.NewTrelloClient(client), nil
}
