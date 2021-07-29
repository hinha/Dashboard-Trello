package repository

import (
	"gorm.io/gorm"
	"time"

	"github.com/hinha/PAM-Trello/app"
)

type trelloRepository struct {
	db *gorm.DB
}

func (r *trelloRepository) Store(in *app.TrelloUserCard) (*app.TrelloUserCard, error) {
	return in, r.db.Create(in).Error
}

func (r *trelloRepository) FindCardCategory(id string) ([]app.CardCategory, error) {
	rows, err := r.db.Raw("select tc.card_category, count(tc.card_category) from trello_user_card tc, trello t "+
		"where tc.card_member_id = t.card_member_id  and t.account_id = ? "+
		"GROUP BY tc.card_category", id).Rows()

	if err != nil {
		return nil, err
	}

	var cards []app.CardCategory

	cards = append(cards, []app.CardCategory{
		{
			Label: string(app.CardTypeTODO),
			Count: 0,
		},
		{
			Label: string(app.CardTypePROGRESS),
			Count: 0,
		},
		{
			Label: string(app.CardTypeDONE),
			Count: 0,
		},
		{
			Label: string(app.CardTypeReview),
			Count: 0,
		},
	}...)

	for rows.Next() {
		var card app.CardCategory
		if err := rows.Scan(&card.Label, &card.Count); err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}

func (r *trelloRepository) CategoryByDate(id string) ([]app.CardGroupBy, error) {
	rows, err := r.db.Raw("select tc.card_category, count(tc.card_member_id) as val, DATE(tc.card_created_at) as date "+
		"from trello_user_card tc, trello t where tc.card_member_id = t.card_member_id and account_id=? group by tc.card_category, "+
		"tc.card_created_at ORDER BY tc.card_created_at ASC", id).Rows()

	if err != nil {
		return nil, err
	}
	layoutDate := "2006-01-02"

	var test []app.CardGroupBy
	for rows.Next() {
		var each app.CardGroupBy
		var date time.Time

		if err := rows.Scan(&each.Category, &each.Count, &date); err != nil {
			return nil, err
		}
		each.Date = date.Format(layoutDate)
		test = append(test, each)
	}

	return unique(test), nil
}

func unique(sample []app.CardGroupBy) []app.CardGroupBy {
	var unique []app.CardGroupBy
sampleLoop:
	for _, v := range sample {
		for i, u := range unique {
			if v.Category == u.Category && v.Date == u.Date {
				unique[i] = v
				unique[i].Count += v.Count
				continue sampleLoop
			}
		}
		unique = append(unique, v)
	}
	return unique
}

func NewTrelloRepository(db *gorm.DB) app.TrelloRepository {
	return &trelloRepository{db: db}
}
