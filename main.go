package main

import (
	"GoodJob/cas"
	"github.com/gorilla/securecookie" // optionally, used for session's encoder/decoder
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
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
		Cookie: cookieName,
		Encode: secureCookie.Encode,
		Decode: secureCookie.Decode,
	})
}

func main() {

	// create our app,
	// set a view
	// set sessions
	// and setup the router for the showcase
	app := iris.New()

	// 可以从任何http相关的恐慌中恢复
	// 并将请求记录到终端。
	app.Use(recover.New())
	goodJobLogger := logger.New()
	app.Use(goodJobLogger)
	app.Use(cas.New(cas.Config{"http://localhost:7007/uuc/login", "http://localhost:3000/a", "http://localhost:7007/uuc"}, sessionsManager))
	//将“before”处理程序注册为将要执行的第一个处理程序
	//在所有域的路由上。
	//或使用`UseGlobal`注册一个将跨子域触发的中间件。
	//app.Use(before)

	//将“after”处理程序注册为将要执行的最后一个处理程序
	//在所有域的路由'处理程序之后。
	//app.Done(after)

	app.OnErrorCode(404, goodJobLogger, func(ctx iris.Context) {
		ctx.Writef("My Custom 404 error page ")
	})

	// or catch all http errors:
	app.OnAnyErrorCode(goodJobLogger, func(ctx iris.Context) {
		// this should be added to the logs, at the end because of the `logger.Config#MessageContextKey`
		ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
		ctx.Writef("My Custom error page")
	})

	app.Get("/a", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// http://localhost:3000
	app.Run(iris.Addr("localhost:3000"), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
}
