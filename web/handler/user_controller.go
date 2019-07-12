package handler

import (
	"github.com/kataras/iris"
	"github.com/lanux/goodjob/v1/user"
)

type UserController struct {
}

func (*UserController) GetReqMapping() string {
	return "/user"
}

func (u *UserController) PartyBuilder(p iris.Party) {
	p.Any("/{id int}", u.GetByUserId)
}

func (*UserController) GetByUserId(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	ctx.StatusCode(iris.StatusOK)
	user, _ := user.GetById(id)
	ctx.JSON(Success(user, ""))
}
