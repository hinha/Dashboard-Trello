package cron_server

import (
	"github.com/gocraft/work"
	log "github.com/sirupsen/logrus"
)

type trelloJobHandler struct {
	//s facebook.Service

	logger *log.Entry
}

func (h *trelloJobHandler) name() string {
	return "trello_stream"
}

func (h *trelloJobHandler) middleware(ctx *work.Job, next work.NextMiddlewareFunc) error {
	return next()
}

func (h *trelloJobHandler) job(c *work.Job) error {
	return nil
}
