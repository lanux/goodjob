package routes

import (
	user2 "goodjob/repositories/rbac"
	user "goodjob/services/rbac"
)

func Get(service user.UserService, id int) (results user2.User) {
	results, _ = service.GetById(id)
	return
}
