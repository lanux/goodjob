package routes

import (
	user2 "GoodJob/repositories/rbac"
	user "GoodJob/services/rbac"
)

func Get(service user.UserService, id int) (results user2.User) {
	results, _ = service.GetById(id)
	return
}
