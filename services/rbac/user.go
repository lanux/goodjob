package user

import "goodjob/repositories/rbac"

type UserService interface {
	GetById(id int) (user.User, bool)
	DeleteById(id int) bool
}

func NewUserService() UserService {
	return &userServiceImpl{
		userRepo: user.NewUserRepository(),
	}
}

type userServiceImpl struct {
	userRepo user.UserRepository
}

func (us *userServiceImpl) GetById(id int) (user.User, bool) {
	return us.userRepo.GetById(id)
}

func (us *userServiceImpl) DeleteById(id int) bool {
	return us.DeleteById(id)
}
