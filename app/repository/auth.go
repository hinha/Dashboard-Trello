package repository

import (
	"time"

	"gorm.io/gorm"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/util/authority"
)

type authRepository struct {
	db *gorm.DB

	access *authority.Authority
}

func (r *authRepository) GetPassword(username string) (*app.Accounts, error) {
	record := new(app.Accounts)
	return record, r.db.Where("username = ?", username).First(record).Error
}

func (r *authRepository) GetRole(userID string) (string, error) {
	// check if the role is a ssigned
	var userRole authority.UserRole
	res := r.db.Where("user_id = ?", userID).First(&userRole)
	if res.Error != nil {
		return "", res.Error
	}

	var role authority.Role
	res = r.db.Where("id = ?", userRole.RoleID).First(&role)
	if res.Error != nil {
		return "", res.Error
	}

	return role.Name, nil
}

func (r *authRepository) UpdateLogin(userID string) error {
	return r.db.Table("accounts").Where("id=?", userID).Updates(app.Accounts{LastLogin: time.Now(), OnlineStatus: true}).Error
}

func NewAuthRepository(db *gorm.DB) app.AuthRepository {
	return &authRepository{db: db, access: authority.New(authority.Options{DB: db, TablesPrefix: "authority_"})}
}
