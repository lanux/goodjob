package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/sessions"
	"github.com/lanux/goodjob/v1/common/logger"
	"github.com/lanux/goodjob/v1/config"
	"github.com/lanux/goodjob/v1/web/handler"
	"github.com/lanux/goodjob/v1/web/middleware/auth"
	"github.com/lanux/goodjob/v1/web/middleware/cas"
	"regexp"
	"strings"
)

// 所有controller实现此接口
type Controller interface {
	RequestMapping() string
	PartyBuilder(p iris.Party)
}

func InitParty(app *iris.Application, sessionsManager *sessions.Sessions) {
	excludes := strings.Split(config.Global.Cas.Excludes, ";")
	casClient := cas.New(&cas.DefaultInterceptor{
		S:        sessionsManager,
		Excludes: &excludes,
		M:        &cas.RegexMatch{Maps: make(map[string]*regexp.Regexp)},
	})
	casbin := auth.New(sessionsManager)
	app.PartyFunc("/api", func(p router.Party) {
		p.Use(casClient.Authentication, casbin.ServeHTTP)
		p.Get("/logout", casClient.RedirectToLogout)
		PartyFunc(p, &handler.UserController{})
	})
}

func PartyFunc(p router.Party, c interface{}) {
	if controller, ok := c.(Controller); ok {
		p.PartyFunc(controller.RequestMapping(), controller.PartyBuilder)
	}
	logger.Warn("not implement Controller")
}
