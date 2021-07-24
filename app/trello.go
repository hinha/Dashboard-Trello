package app

import "time"

type TrelloRepository interface {
	Store(in *TrelloUserCard) (*TrelloUserCard, error)
}

type Trello struct {
	ID         uint64 `json:"id" gorm:"primary_key"`
	ApiKey     string `json:"api_key" gorm:"type:text;not null"`
	Token      string `json:"token" gorm:"type:text;not null"`
	BoardName  string `json:"board_name" gorm:"type:varchar(50);not null"`
	BoardID    string `json:"board_id" gorm:"type:varchar(50);not null"`
	CardUserID string `json:"card_user_id" gorm:"type:varchar(120);not null"`
	AccountID  string `json:"-" gorm:"type:varchar(50)"`
}

type TrelloUserCard struct {
	ID                       string    `json:"id" gorm:"type:varchar(120);primaryKey"`
	CardID                   string    `json:"card_id" gorm:"type:varchar(120);not null"`
	CardName                 string    `json:"card_name" gorm:"type:text;not null"`
	CardCategory             string    `json:"card_category" gorm:"type:varchar(120);not null"`
	CardVotes                int64     `json:"card_votes" gorm:"type:int(12);not null"`
	CardCheckItems           int64     `json:"card_check_items" gorm:"type:int(12);not null"`
	CardCheckLists           int64     `json:"card_check_lists" gorm:"type:int(12);not null"`
	CardCommentCount         int64     `json:"comment_count" gorm:"type:int(12);not null"`
	CardAttachmentsCount     int64     `json:"attachments_count" gorm:"type:int(12);not null"`
	CardCheckItemsComplete   int64     `json:"card_check_items_complete" gorm:"type:int(12);not null"`
	CardCheckItemsInComplete int64     `json:"card_check_items_incomplete" gorm:"type:int(12);not null"`
	CardMemberName           string    `json:"card_member_name" gorm:"type:text;not null"`
	CardMemberUsername       string    `json:"card_member_username" gorm:"type:text;not null"`
	CardCreatedAt            time.Time `json:"card_created_at" gorm:"not null;"`
	Init                     Trello    `json:"-" gorm:"ForeignKey:CardUserID"`
}
