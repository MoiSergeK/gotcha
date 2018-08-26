package main

import (
	"github.com/kataras/iris"
	// "github.com/kataras/iris/middleware/logger"
	// "github.com/kataras/iris/middleware/recover"
	socketio "github.com/googollee/go-socket.io"
	"./app"
	"./app/db"
	"time"
)

func main() {
	defer db.DBConn().Close()
	
	server := iris.New()

	io, err := socketio.NewServer(nil)

	app.InitSocketIO(io)
	
	if err != nil { panic(err) }
	
	server.Logger().SetLevel("debug")
	
	// server.Use(recover.New())
	// server.Use(logger.New())
	
	app.Router(server)
	
	db.SyncDB()

	server.Any("/socket.io/{p:path}", iris.FromStd(io))

	// server.StaticWeb("/", "./app/views")

	server.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func cron() {
	for ;; {
		time.Sleep(time.Second * 5)
	}
}