package accounts

import (
	"context"
	"fmt"

	"github.com/hinha/PAM-Trello/app"
)

type Service interface {
	AuthLogin(ctx context.Context, in *app.LoginInput) (*app.Accounts, string, error)
}

type service struct {
	auth app.AuthRepository
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

func NewService(auth app.AuthRepository) *service {
	return &service{auth: auth}
}
