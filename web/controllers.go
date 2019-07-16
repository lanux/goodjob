package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/lanux/goodjob/v1/common/logger"
	"github.com/lanux/goodjob/v1/config"
	"github.com/lanux/goodjob/v1/web/handler"
	"github.com/lanux/goodjob/v1/web/middleware/auth"
	"github.com/lanux/goodjob/v1/web/middleware/cas"
	"regexp"
	"strings"
)

// 所有controller实现此接口
type controller interface {
	RequestMapping() string
	PartyBuilder(p iris.Party)
}

func InitParty(app *iris.Application) {
	excludes := strings.Split(config.Cas.Excludes, ";")
	casClient := cas.New(&cas.DefaultInterceptor{
		S:        Sessions,
		Excludes: &excludes,
		M:        &cas.RegexMatch{Maps: make(map[string]*regexp.Regexp)},
	})
	casbin := auth.New(Sessions)
	app.PartyFunc("/api", func(p router.Party) {
		p.Use(casClient.Authentication, casbin.ServeHTTP)
		p.Get("/logout", casClient.RedirectToLogout)
		initPartyFunc(p, &handler.UserController{})
	})
}

func initPartyFunc(p router.Party, c controller) {
	logger.Info("init party path:" + c.RequestMapping())
	p.PartyFunc(c.RequestMapping(), c.PartyBuilder)
}

func partyFunc(p router.Party, c interface{}) {
	if controller, ok := c.(controller); ok {
		p.PartyFunc(controller.RequestMapping(), controller.PartyBuilder)
	}
	logger.Warn("not implement Controller")
}
