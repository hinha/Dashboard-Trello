package authority

// UserRole represents the relationship between users and roles
type UserRole struct {
	ID     uint
	UserID string `json:"user_id" gorm:"type:varchar(50);not null;ForeignKey:ID"`
	RoleID uint
}

// TableName sets the table name
func (u UserRole) TableName() string {
	return TablePrefix + "user_roles"
}
