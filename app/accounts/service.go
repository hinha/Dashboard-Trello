package accounts

import (
	"context"
	"fmt"

	"github.com/hinha/PAM-Trello/app"
)

type Service interface {
	AuthLogin(ctx context.Context, in *app.LoginInput) (*app.Accounts, string, error)
	NewAccount(ctx context.Context, adminID string, roleName string, in *app.RegisterInput) error
	ListAccount(ctx context.Context, adminId string, roleName string) ([]app.Accounts, error)
}

type service struct {
	auth    app.AuthRepository
	account app.AccountRepository
}

func (s *service) AuthLogin(ctx context.Context, in *app.LoginInput) (*app.Accounts, string, error) {
	account, err := s.auth.GetPassword(in.Username)
	if err != nil {
		return nil, "", fmt.Errorf("username or password not valid")
	}

	if !in.ComparePassword(account.Password, account.SecretPassword) {
		return nil, "", fmt.Errorf("username or password not valid")
	}

	role, err := s.auth.GetRole(account.ID)
	if err != nil {
		return nil, "", fmt.Errorf("access denied")
	}

	data := map[string]interface{}{
		"status": role,
		"time":   account.LastLogin.Unix(),
	}

	token, err := in.GenerateJwt(data)
	if err != nil {
		return nil, "", fmt.Errorf("service unavailable")
	}

	_ = s.auth.UpdateLogin(account.ID)

	return account, token, nil
}

func (s *service) NewAccount(ctx context.Context, adminID string, roleName string, in *app.RegisterInput) error {

	record, err := s.account.FindUsername(in.Username)
	if err != nil {
		return fmt.Errorf("error when inserted data")
	}

	if record.Username == "" {
		err := s.account.Store(adminID, roleName, in)
		if err != nil {
			return err
		}

		// it success register data
		return nil
	}
	return fmt.Errorf("user already registered")
}

func (s *service) ListAccount(ctx context.Context, adminId string, roleName string) ([]app.Accounts, error) {
	return s.account.GetAccount(adminId, roleName)
}

func NewService(auth app.AuthRepository, account app.AccountRepository) *service {
	return &service{auth: auth, account: account}
}
