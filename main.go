package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"./app"
)

func main() {
	server := iris.New()

	server.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	server.Use(recover.New())
	server.Use(logger.New())

	app.Router(server)

	server.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}