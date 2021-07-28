package app

import (
	"gorm.io/gorm"
	"time"
)

type TrelloRepository interface {
	Store(in *TrelloUserCard) (*TrelloUserCard, error)
	FindCardCategory(id string, category string) (int, error)
}

type Trello struct {
	ID           uint64           `json:"id" gorm:"primary_key"`
	BoardName    string           `json:"board_name" gorm:"type:varchar(50);not null"`
	BoardID      string           `json:"board_id" gorm:"type:varchar(50);not null"`
	CardMemberID string           `json:"-" gorm:"type:varchar(120);null;index:card_member_id_id_k"`
	AccountID    string           `json:"-" gorm:"type:varchar(50);null"`
	CardItems    []TrelloUserCard `json:"card_items" gorm:"ForeignKey:CardMemberID;references:CardMemberID"`
}

func (Trello) TableName() string {
	return "trello"
}

type TrelloUserCard struct {
	gorm.Model
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
	CardUrl                  string    `json:"card_url" gorm:"type:text;not null"`
	CardMemberID             string    `json:"card_member_id" gorm:"type:varchar(120);not null"`
	CardMemberName           string    `json:"card_member_name" gorm:"type:text;not null"`
	CardMemberUsername       string    `json:"card_member_username" gorm:"type:text;not null"`
	CardCreatedAt            time.Time `json:"card_created_at" gorm:"not null;"`
}

func (TrelloUserCard) TableName() string {
	return "trello_user_card"
}

type Performance struct {
	Todo          int         `json:"todo"`
	OnProgress    int         `json:"on_progress"`
	Done          int         `json:"done"`
	Product       interface{} `json:"product"`
	Daily         interface{} `json:"daily"`
	TimelineGantt interface{} `json:"timeline_gantt"`
}
