package app

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type AccountRepository interface {
	Store(adminID, roleName string, in *RegisterInput) error
	FindUsername(username string) (*Accounts, error)
	FindID(id string) (*Accounts, error)
	GetAccount(adminID string, roleName string) ([]Accounts, error)
	DeleteAccount(id string, username string) error
	CheckRole(adminID string, roleName string) error
	GivenPermission(userId string, roleName string, permName string) error
	AssignAccessControl(adminID string, roleName string, control *AssignRole) error
	AccessControlList() (AccessControl, error)
	SetOnlineStatus(id string, status bool) error
	GetOnlineStatus(Id string) ([]*Accounts, error)
	ListAccount(ignoreID string) ([]*Accounts, error)
}

type AccountDetail struct {
	ID            string    `json:"id" gorm:"type:varchar(50);primaryKey"`
	JobTitle      string    `json:"job_title" gorm:"type:text;not null"`
	Level         string    `json:"level" gorm:"type:text;not null"`
	Phone         string    `json:"phone" gorm:"type:text;not null"`
	Address       string    `json:"address" gorm:"type:text;not null"`
	StartContract time.Time `json:"start_contract" gorm:"not null;"`
	EndContract   time.Time `json:"end_contract" gorm:"not null;"`
	AccountID     string    `json:"-" gorm:"type:varchar(50)"`
}

type RegisterInput struct {
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Errors   map[string]string `json:"-"`
	Token    string
}

func (m *RegisterInput) ValidateAPI() bool {
	m.Errors = make(map[string]string)

	if strings.TrimSpace(m.Name) == "" {
		m.Errors["message"] = "Please enter a name"
	} else if strings.TrimSpace(m.Email) == "" {
		m.Errors["message"] = "Please enter a email"
	} else if strings.TrimSpace(m.Username) == "" {
		m.Errors["message"] = "Please enter a username"
	} else if strings.TrimSpace(m.Password) == "" {
		m.Errors["message"] = "Please enter a password"
	}
	return len(m.Errors) == 0
}

func (m *RegisterInput) GeneratePassword(secret string) (string, error) {
	password := []byte(fmt.Sprintf("%s:%s", secret, m.Password))
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

type AccessControl struct {
	Authority []struct {
		Name string `json:"name"`
	} `json:"authority"`
	Permission []struct {
		Name string `json:"name"`
	} `json:"permission"`
}

type AssignRole struct {
	UserID     string            `json:"userID"`
	Role       string            `json:"role"`
	Permission string            `json:"permission"`
	Errors     map[string]string `json:"-"`
}

func (m *AssignRole) Validate() bool {
	m.Errors = make(map[string]string)

	if strings.TrimSpace(m.UserID) == "" {
		m.Errors["message"] = "Empty user"
	} else if strings.TrimSpace(m.Role) == "" {
		m.Errors["message"] = "Please select a role"
	} else if strings.TrimSpace(m.Permission) == "" {
		m.Errors["message"] = "Please select a permission"
	}

	return len(m.Errors) == 0
}
