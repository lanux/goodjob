package handler

import (
	"github.com/kataras/iris"
	"github.com/lanux/goodjob/v1/service"
)

type UserController struct {
}

var userService = new(service.UserService)

func (*UserController) RequestMapping() string {
	return "/user"
}

func (u *UserController) PartyBuilder(p iris.Party) {
	p.Get("/{id int}", u.GetByUserId)
}

func (*UserController) GetByUserId(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	ctx.StatusCode(iris.StatusOK)
	user := userService.GetById(id)
	ctx.JSON(Success(user, ""))
}
