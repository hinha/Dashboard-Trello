package app

import (
	"database/sql/driver"
	"time"
)

type AttendanceActivity string

const (
	AttendanceIN  AttendanceActivity = "IN"
	AttendanceOUT AttendanceActivity = "OUT"
)

func (e *AttendanceActivity) Scan(value interface{}) error {
	*e = AttendanceActivity(value.([]byte))
	return nil
}

func (e AttendanceActivity) Value() (driver.Value, error) {
	return string(e), nil
}

type AttendanceLabel string

const (
	AttendanceMeeting AttendanceLabel = "MEETING"
	AttendanceWork    AttendanceLabel = "WORK"
	AttendanceBreak   AttendanceLabel = "BREAK"
)

func (e *AttendanceLabel) Scan(value interface{}) error {
	*e = AttendanceLabel(value.([]byte))
	return nil
}

func (e AttendanceLabel) Value() (driver.Value, error) {
	return string(e), nil
}

type Attendance struct {
	ID             string             `json:"id" gorm:"type:varchar(50);primaryKey"`
	ActivityStatus AttendanceActivity `json:"activity_status" gorm:"type:enum('IN', 'OUT');default:'IN'"`
	Date           time.Time          `json:"date" gorm:"not null;"`
	Description    string             `json:"description" gorm:"type:text;null"`
	Label          AttendanceLabel    `json:"label" gorm:"type:enum('MEETING', 'WORK', 'BREAK');default:'WORK'"`
	Status         int                `json:"status" gorm:"type:int(10);default:0"`
	ActualSignIn   time.Time          `json:"actual_sign_in" gorm:"not null;"`
	ActualSignOut  time.Time          `json:"actual_sign_out" gorm:"not null;"`
	LateTime       int                `json:"late_time" gorm:"type:int(10);default:0"`

	UpdatedAt time.Time `json:"-"`
	AccountID string
}
