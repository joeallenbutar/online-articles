package entity

type User struct {
	Username string `gorm:"primaryKey;column:username;type:varchar(100)"`
	Password string `gorm:"column:password;type:varchar(200)"`
	IsActive int    `gorm:"column:is_active;default:1"`
}

func (User) TableName() string {
	return "tb_user"
}
