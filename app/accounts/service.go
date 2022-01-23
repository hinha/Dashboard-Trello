package accounts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/util/security"
)

type Service interface {
	Authorize(key string) (interface{}, error)
	AuthLogin(ctx context.Context, in *app.LoginInput) (*app.Accounts, string, error)
	GetProfile(ctx context.Context, id string) (*app.Accounts, []string, string, error)
	NewAccount(ctx context.Context, adminID string, roleName string, in *app.RegisterInput) (*app.Accounts, error)
	ListAccount(ctx context.Context, adminId string, roleName string) ([]app.Accounts, error)
	DeleteAccount(ctx context.Context, adminId string, roleName string, userID string, userName string) error
	UpdateAccount(ctx context.Context, adminId string, roleName string, account app.UpdateAccount) error
	GetAccessList(ctx context.Context) (app.AccessControl, error)
	NewAccessControlList(ctx context.Context, adminId string, roleAdmin string, control *app.AssignRole) error
	GetDetailAccount(ctx context.Context, accountID string) (*app.AccountDetail, error)
}

type service struct {
	auth    app.AuthRepository
	account app.AccountRepository

	encrypt *security.BearerCipher
}

// Authorize Bearer cipher Cbc
func (s *service) Authorize(key string) (interface{}, error) {
	return security.Authorize(s.encrypt, key)
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
		"id":     account.ID,
	}

	token, err := in.GenerateJwt(data)
	if err != nil {
		return nil, "", fmt.Errorf("service unavailable")
	}

	_ = s.auth.UpdateLogin(account.ID)

	return account, token, nil
}

func (s *service) NewAccount(ctx context.Context, adminID string, roleName string, in *app.RegisterInput) (*app.Accounts, error) {

	record, err := s.account.FindUsername(in.Username)
	if err != nil {
		return nil, fmt.Errorf("error when inserted data")
	}

	if record.Username == "" {
		err := s.account.Store(adminID, roleName, in)
		if err != nil {
			return nil, err
		}

		// it success register data
		return s.account.FindUsername(in.Username)
	}
	return nil, fmt.Errorf("user already registered")
}

func (s *service) ListAccount(ctx context.Context, adminId string, roleName string) ([]app.Accounts, error) {
	return s.account.GetAccount(adminId, roleName)
}

func (s *service) DeleteAccount(ctx context.Context, adminId string, roleName string, userID string, userName string) error {
	if err := s.account.CheckRole(adminId, roleName); err != nil {
		return err
	}

	if err := s.account.DeleteAccount(userID, userName); err != nil {
		return fmt.Errorf("error when delete")
	}

	return nil
}

func (s *service) UpdateAccount(ctx context.Context, adminId string, roleName string, account app.UpdateAccount) error {
	if err := s.account.CheckRole(adminId, roleName); err != nil {
		return err
	}

	if err := s.account.UpdateAccount(account); err != nil {
		return err
	}

	return nil
}

func (s *service) GetAccessList(ctx context.Context) (app.AccessControl, error) {
	control, err := s.account.AccessControlList()
	if err != nil {
		return control, err
	}

	return control, nil
}

func (s *service) NewAccessControlList(ctx context.Context, adminId string, roleAdmin string, control *app.AssignRole) error {
	if err := s.account.GivenPermission(control.UserID, control.Role, control.Permission); err != nil {
		if err.Error() == "user doesn't have a role assigned" {
			return s.account.AssignAccessControl(adminId, roleAdmin, control)
		}
		return err
	}

	return fmt.Errorf("something error when assign permission")
}

// GetProfile ...
// return value: app.Accounts, []resource_permission, secret payload, error
func (s *service) GetProfile(ctx context.Context, id string) (*app.Accounts, []string, string, error) {
	account, err := s.account.FindID(id)
	if err != nil {
		return account, nil, "", err
	}

	role, err := s.auth.GetRole(account.ID)
	if err != nil {
		return nil, nil, "", fmt.Errorf("access denied")
	}

	payload := account.Payload(role)
	bytes, _ := json.Marshal(payload)

	secret, err := s.encrypt.EncryptStringCBC(bytes)
	if err != nil {
		return account, nil, "", err
	}
	return account, account.ResourcePermission(role), secret, nil
}

func (s *service) GetDetailAccount(ctx context.Context, accountID string) (*app.AccountDetail, error) {
	account, err := s.account.GetAccountDetail(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func NewService(auth app.AuthRepository, account app.AccountRepository) *service {
	CipherIv := "Programmer is not robot"
	CipherHeader := "sangatrahasiabro[HEHE]"
	CipherKey := "Harga kopi ditentukan oleh kualitas"
	return &service{auth: auth, account: account, encrypt: security.NewCipher(CipherIv, CipherHeader, CipherKey)}
}
