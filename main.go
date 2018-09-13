package main

import (
	"github.com/kataras/iris"
	"./app"
	"./app/db"
	// "./app/libs"
	"time"
)

func main() {
	defer db.DBConn().Close()
	
	server := iris.New()
	
	server.Logger().SetLevel("debug")

	// libs.LoadConfig()
	
	app.Router(server)
	
	db.SyncDB()

	server.StaticWeb("/", "./app/views")

	server.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func cron() {
	for ;; {
		time.Sleep(time.Second * 5)
	}
}