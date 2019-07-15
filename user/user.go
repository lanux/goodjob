package user

import (
	"fmt"
	"github.com/lanux/goodjob/v1/db"
)

type User struct {
	Id       int   `gorm:"primary_key"`
	Name     string `gorm:"not null VARCHAR(191)"`
	Account  string `gorm:"unique;VARCHAR(191)"`
	Password string `gorm:"not null VARCHAR(191)"`
}

func (User) TableName() string {
	return "sys_user"
}

func Create(user *User) bool {
	ok := db.Instance().NewRecord(user)
	fmt.Printf("create user %s", ok)
	return ok
}

func GetById(id int) (User, bool) {
	user := new(User)
	user.Id = id
	if err := db.Instance().Preload("id").First(user).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
	}
	return *user, true
}

func DeleteById(id int) bool {
	u := new(User)
	u.Id = id
	if err := db.Instance().Delete(u).Error; err != nil {
		fmt.Printf("DeleteUserByIdErr:%s", err)
	}
	return true
}
