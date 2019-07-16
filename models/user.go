package models

type User struct {
	Id       int   `gorm:"primary_key"`
	Name     string `gorm:"not null VARCHAR(191)"`
	Account  string `gorm:"unique;VARCHAR(191)"`
	Password string `gorm:"not null VARCHAR(191)"`
}

func (User) TableName() string {
	return "sys_user"
}

