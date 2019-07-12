package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/lanux/goodjob/v1/web/handler"
	"github.com/lanux/goodjob/v1/web/middleware/cas"
)

type Controller interface {
	GetReqMapping() string
	PartyBuilder(p iris.Party)
}

func InitParty(app *iris.Application, sessionsManager *sessions.Sessions) {
	app.Get("/", func(context context.Context) {
		session := sessionsManager.Start(context)
		user := session.Get("user").(cas.AuthSuccessStruct)
		context.Markdown([]byte("### HELLO " + user.Attributes.UserName))
	})
	app.Get("/logout", cas.C.RedirectToLogout)
	userController := &handler.UserController{}
	app.PartyFunc(userController.GetReqMapping(), userController.PartyBuilder)
}
