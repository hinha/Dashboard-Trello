package repository

import (
	"fmt"
	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/util/authority"
	"github.com/hinha/PAM-Trello/app/util/generate"
	"gorm.io/gorm"
	"math"
	"time"
)

type accountRepository struct {
	db *gorm.DB

	access *authority.Authority
}

func (r *accountRepository) Store(adminID, roleName string, in *app.RegisterInput) error {
	ok, err := r.access.CheckRole(adminID, roleName)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("access permission denied, please contact admin")
	}

	secretPassword, _ := generate.GenerateRandomString(24)

	password, err := in.GeneratePassword(secretPassword)
	if err != nil {
		return err
	}

	UserID := fmt.Sprintf("%d%d", len(in.Username), int(time.Now().Unix()/10)%math.MaxInt64)
	create := r.db.Create(&app.Accounts{
		ID:             UserID,
		Name:           in.Name,
		Username:       in.Username,
		Password:       password,
		SecretPassword: secretPassword,
		Email:          in.Email,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastLogin:      time.Now(),
	})

	return create.Error
}

func (r *accountRepository) FindUsername(username string) (*app.Accounts, error) {
	account := new(app.Accounts)
	err := r.db.Find(account, "username = ?", username).Error
	return account, err
}

func (r *accountRepository) GetAccount(adminID string, roleName string) ([]app.Accounts, error) {
	ok, err := r.access.CheckRole(adminID, roleName)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("access permission denied, please contact admin")
	}

	var accounts []app.Accounts
	err = r.db.Find(&accounts).Limit(10).Error

	return accounts, err
}

func NewAccountRepository(db *gorm.DB) app.AccountRepository {
	return &accountRepository{db: db, access: authority.New(authority.Options{DB: db, TablesPrefix: "authority_"})}
}
