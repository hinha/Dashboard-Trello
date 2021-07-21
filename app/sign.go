package app

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	GetPassword(username string) (*Accounts, error)
	GetRole(userID string) (string, error)
	UpdateLogin(userID string) error
}

type Accounts struct {
	ID             string    `json:"id" gorm:"type:varchar(50);primaryKey"`
	Name           string    `json:"name" gorm:"type:varchar(50);not null"`
	Username       string    `json:"username" gorm:"type:varchar(24);not null"`
	Password       string    `json:"-" gorm:"type:text;not null"`
	SecretPassword string    `json:"-" gorm:"type:text;not null"`
	Email          string    `json:"email" gorm:"type:varchar(80);not null"`
	OnlineStatus   bool      `json:"online_status" gorm:"default:false"`
	SuspendStatus  bool      `json:"suspend_status" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null;"`
	UpdatedAt      time.Time
	LastLogin      time.Time     `json:"last_login" gorm:"not null;"`
	Accounts       AccountDetail `json:"-" gorm:"ForeignKey:AccountID"`
}

type LoginInput struct {
	Username string
	Password string
	Errors   map[string]string
	Token    string
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

func (m *LoginInput) RefreshJwt(oldToken string) (string, error) {
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims
	if err := claims.Valid(); err != nil {
		return "", err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	for k, v := range claims.(jwt.MapClaims) {
		rtClaims[k] = v
	}
	rtClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	rt, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return rt, nil
}

const (
	PermUserGroup          = "add users and groups"
	PermUserProperties     = "manage most users properties"
	PermUserManage         = "create and manage user views"
	PermUserUpdatePassword = "update password expiration policies"
	PermServiceRequest     = "manage service requests"
	PermReadWrite          = "readWrite"
	PermResetPassword      = "resetPassword"
	PermListEmployee       = "listEmployee"
	PermAttendance         = "readWriteAttendance"
	PermUserDetails        = "readUpdateUserDetails"
)

var DeleteCookie = &http.Cookie{
	Name:    "token",
	Value:   "",
	Path:    "/",
	Expires: time.Now().Add(0),
}
