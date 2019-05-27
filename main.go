package main

import (
	"github.com/casbin/casbin"
	"github.com/gorilla/securecookie" // optionally, used for session's encoder/decoder
	cm "github.com/iris-contrib/middleware/casbin"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/sessions"
	"goodjob/config"
	"goodjob/db"
	"goodjob/services/rbac"
	"goodjob/web/middleware"
	"goodjob/web/routes"
	"goodjob/web/structs"
	"runtime"
	"time"
)

var sessionsManager *sessions.Sessions

func init() {
	// attach a session manager
	cookieName := "GOSESSIONID"
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie := securecookie.New(hashKey, blockKey)

	sessionsManager = sessions.New(sessions.Config{
		Cookie:  cookieName,
		Encode:  secureCookie.Encode,
		Decode:  secureCookie.Decode,
		Expires: time.Duration(30) * time.Minute,
	})
}

func main() {
	// Set the concurrency level
	runtime.GOMAXPROCS(4 * runtime.NumCPU())

	app := iris.Default()

	//app.Logger().AddOutput(f) 指定输出到文件
	app.Logger().SetLevel("debug")

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() {
		db.MYSQL.Close()
	})

	app.Use(cas.New(func(ctx iris.Context) bool {
		session := sessionsManager.Start(ctx)
		return session.Get("user") == nil
	}, func(ctx iris.Context, u interface{}) {
		if u != nil {
			session := sessionsManager.Start(ctx)
			session.Set("user", u)
		}
	}))

	globalLocale := i18n.New(i18n.Config{
		Default:      "zh-CN",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "./locales/locale_en-US.ini",
			"zh-CN": "./locales/locale_zh-CN.ini"}})
	app.Use(globalLocale)

	enforcer := casbin.NewEnforcer("./config/casbinmodel.conf", "casbinpolicy.csv")
	enforcer.EnableLog(true)
	casbinMiddleware := cm.New(enforcer)
	app.Use(casbinMiddleware.ServeHTTP)
	app.WrapRouter(casbinMiddleware.Wrapper())

	//将“before”处理程序注册为将要执行的第一个处理程序
	//在所有域的路由上。
	//或使用`UseGlobal`注册一个将跨子域触发的中间件。
	//app.Use(before)

	//将“after”处理程序注册为将要执行的最后一个处理程序
	//在所有域的路由'处理程序之后。
	//app.Done(after)

	// or catch all http errors:
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.JSON(structs.ResponseBasic{
			Code:    ctx.GetStatusCode(),
			Message: i18n.Translate(ctx, "error")})
	})
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(structs.ResponseBasic{
			Code:    ctx.GetStatusCode(),
			Message: i18n.Translate(ctx, "404")})
	})

	app.Favicon("./assets/logo_24.ico")

	hero.Register(user.NewUserService())
	app.PartyFunc("/user", func(r iris.Party) {
		r.Get("/{id:int}", hero.Handler(routes.Get))
	})

	app.Run(iris.Addr(*config.Global.Host+":"+*config.Global.Port), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
	//app.Run(iris.AutoTLS(":443", "example.com", "admin@example.com"))

}
