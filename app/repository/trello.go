package repository

import (
	"github.com/hinha/PAM-Trello/app"
	"gorm.io/gorm"
)

type trelloRepository struct {
	db *gorm.DB
}

func (r *trelloRepository) Store(in *app.TrelloUserCard) (*app.TrelloUserCard, error) {
	return in, r.db.Create(in).Error
}

func NewTrelloRepository(db *gorm.DB) app.TrelloRepository {
	return &trelloRepository{db: db}
}
