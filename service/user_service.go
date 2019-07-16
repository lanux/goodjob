package service

import (
	"github.com/lanux/goodjob/v1/dao"
	"github.com/lanux/goodjob/v1/models"
)

//UserService
type UserService struct {
}

var userDao = new(dao.UserDao)

func (*UserService) GetById(id int64) *models.User {
	u := userDao.GetById(id, &models.User{})
	return u.(*models.User)
}
