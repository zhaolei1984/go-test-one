package iris

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func MyHello(ctx context.Context) {
	ctx.ViewData("message", "Hello world!")
	ctx.View("hello.html")
}

func RunIrisServe()  {
	app := iris.Default()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", MyHello)

	app.Run(iris.Addr(":7000"))
}