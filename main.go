package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/lanux/goodjob/v1/config"
	"github.com/lanux/goodjob/v1/db"
	"github.com/lanux/goodjob/v1/web"
	"github.com/lanux/goodjob/v1/web/handler"
	"runtime"
)

func main() {
	// Set the concurrency level
	runtime.GOMAXPROCS(4 * runtime.NumCPU())

	app := iris.Default()

	//app.Logger().AddOutput(f) 指定输出到文件
	app.Logger().SetLevel("debug")

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() {
		db.Instance().Close()
	})

	globalLocale := i18n.New(i18n.Config{
		Default:      "zh-CN",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "./locales/locale_en-US.ini",
			"zh-CN": "./locales/locale_zh-CN.ini"}})
	app.Use(globalLocale)

	//将“before”处理程序注册为将要执行的第一个处理程序
	//在所有域的路由上。
	//或使用`UseGlobal`注册一个将跨子域触发的中间件。
	//app.Use(before)

	//将“after”处理程序注册为将要执行的最后一个处理程序
	//在所有域的路由'处理程序之后。
	//app.Done(after)

	// or catch all http errors:
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.JSON(handler.ErrorWithLocale(ctx.GetStatusCode(), "error", ctx))
	})
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(handler.ErrorWithLocale(ctx.GetStatusCode(), "404", ctx))
	})

	app.Favicon("./assets/logo_24.ico")

	web.InitParty(app)

	app.Run(iris.Addr(config.Host+":"+config.Port), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
	//app.Run(iris.AutoTLS(":443", "example.com", "admin@example.com"))
}
