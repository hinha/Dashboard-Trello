package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	GetPassword(username string) (*Accounts, error)
	GetRole(userID string) (string, error)
}

type Accounts struct {
	ID             string    `json:"id" gorm:"type:varchar(50);primaryKey"`
	Name           string    `json:"username" gorm:"type:varchar(50);not null"`
	Username       string    `json:"username" gorm:"type:varchar(24);not null"`
	Password       string    `json:"-" gorm:"type:text;not null"`
	SecretPassword string    `json:"-" gorm:"type:text;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null;"`
	LastLogin      time.Time `json:"last_login" gorm:"not null;"`
	//RoleID string `gorm:"-"`
}

type LoginInput struct {
	Username string
	Password string
	Errors   map[string]string
}

func (m *LoginInput) Validate() bool {
	m.Errors = make(map[string]string)

	if strings.TrimSpace(m.Username) == "" {
		m.Errors["Username"] = "Please enter a username"
	} else if strings.TrimSpace(m.Password) == "" {
		m.Errors["Password"] = "Please enter a password"
	}
	return len(m.Errors) == 0
}

func (m *LoginInput) ComparePassword(hashed string, secret string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(fmt.Sprintf("%s:%s", secret, m.Password)))
	if err != nil {
		return false
	}

	return true
}

func (m *LoginInput) GenerateJwt(data map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	for k, v := range data {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

const (
	PermUserGroup          = "add users and groups"
	PermUserProperties     = "manage most users properties"
	PermUserManage         = "create and manage user views"
	PermUserUpdatePassword = "update password expiration policies"
	PermServiceRequest     = "manage service requests"
)
