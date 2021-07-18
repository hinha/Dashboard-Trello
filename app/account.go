package app

import "time"

type AccountDetail struct {
	ID            string    `json:"id" gorm:"type:varchar(50);primaryKey"`
	Email         string    `json:"email" gorm:"type:varchar(80);not null"`
	JobTitle      string    `json:"job_title" gorm:"type:text;not null"`
	Level         string    `json:"level" gorm:"type:text;not null"`
	Phone         string    `json:"phone" gorm:"type:text;not null"`
	Address       string    `json:"address" gorm:"type:text;not null"`
	StartContract time.Time `json:"start_contract" gorm:"not null;"`
	EndContract   time.Time `json:"end_contract" gorm:"not null;"`
	AccountID     string    `json:"-" gorm:"type:varchar(50)"`
}
