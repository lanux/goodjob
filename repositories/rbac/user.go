package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goodjob/db"
)

type User struct {
	gorm.Model
	Id       int    `gorm:"not null int(11)"`
	Name     string `gorm:"not null VARCHAR(191)"`
	Username string `gorm:"unique;VARCHAR(191)"`
	Password string `gorm:"not null VARCHAR(191)"`
	RoleID   uint
}

type UserRepository interface {
	GetById(id int) (User, bool)
	DeleteById(id int) bool
}

func NewUserRepository() UserRepository {
	return &userDbRepository{}
}

type userDbRepository struct {
}

func (*userDbRepository) GetById(id int) (User, bool) {
	user := new(User)
	user.Id = id
	if err := db.MYSQL.Preload("Role").First(user).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
	}
	return *user, true
}

func (*userDbRepository) DeleteById(id int) bool {
	u := new(User)
	u.Id = id
	if err := db.MYSQL.Delete(u).Error; err != nil {
		fmt.Printf("DeleteUserByIdErr:%s", err)
	}
	return true
}
