package repository

import (
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/util/authority"
	"github.com/hinha/PAM-Trello/app/util/generate"
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
	err = r.db.Find(&accounts).Order("created_at desc").Limit(10).Error

	return accounts, err
}

func (r *accountRepository) CheckRole(adminID string, roleName string) error {
	ok, err := r.access.CheckRole(adminID, roleName)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("access permission denied, please contact admin")
	}

	return nil
}

func (r *accountRepository) DeleteAccount(id string, username string) error {
	return r.db.Where("id = ? and username = ?", id, username).Delete(app.Accounts{}).Error
}

func (r accountRepository) GivenPermission(userId string, roleName, permName string) error {
	// check if a role have a given permission
	ok, err := r.access.CheckRolePermission(roleName, permName)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("permission not allowed by role name")
	}

	ok, err = r.access.CheckPermission(userId, permName)
	if err != nil {
		return err
	}

	if ok {
		return fmt.Errorf("have already registered permission access")
	}

	return nil
}

func (r *accountRepository) AssignAccessControl(adminID string, roleName string, control *app.AssignRole) error {
	ok, err := r.access.CheckRole(adminID, roleName)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("access permission denied, please contact admin")
	}

	// check if a role have a given permission
	ok, err = r.access.CheckRolePermission(control.Role, control.Permission)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("role not defined from registry")
	}

	// allowed status if true
	if err := r.access.AssignRole(control.UserID, control.Role); err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) AccessControlList() (app.AccessControl, error) {
	var result app.AccessControl

	if err := r.db.Table("authority_roles").Find(&result.Authority).Error; err != nil {
		return result, err
	}
	if err := r.db.Table("authority_permissions").Find(&result.Permission).Error; err != nil {
		return result, err
	}

	return result, nil
}

func NewAccountRepository(db *gorm.DB) app.AccountRepository {
	return &accountRepository{db: db, access: authority.New(authority.Options{DB: db, TablesPrefix: "authority_"})}
}
