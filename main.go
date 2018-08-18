package main

import (
	"github.com/kataras/iris"
	// "github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"./app"
	"./app/libs"
)

func main() {
	defer libs.DBConn().Close()
	
	server := iris.New()
	
	server.Logger().SetLevel("debug")
	
	server.Use(recover.New())
	// server.Use(logger.New())
	
	app.Router(server)
	
	libs.SyncDB()

	server.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}