package repository

import (
	"gorm.io/gorm"

	"github.com/hinha/PAM-Trello/app"
)

type trelloRepository struct {
	db *gorm.DB
}

func (r *trelloRepository) Store(in *app.TrelloUserCard) (*app.TrelloUserCard, error) {
	return in, r.db.Create(in).Error
}

func (r *trelloRepository) FindCardCategory(id string, category string) (int, error) {
	var count int

	err := r.db.Raw("select count(tc.card_category) from trello_user_card tc, trello t "+
		"where tc.card_member_id = t.card_member_id  and t.account_id = ? "+
		"GROUP BY tc.card_category HAVING tc.card_category = ?", id, category).Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewTrelloRepository(db *gorm.DB) app.TrelloRepository {
	return &trelloRepository{db: db}
}
