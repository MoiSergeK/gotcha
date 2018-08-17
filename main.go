package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"./app"
	"./app/libs"
)

func main() {
	defer libs.DBConn().Close()
	
	server := iris.New()
	server.Logger().SetLevel("debug")
	server.Use(recover.New())
	server.Use(logger.New())
	app.Router(server)
	server.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

	// libs.Insert("places", map[string]string{"coords": "123.12,12.3", "address": "bbbbbb"})
	// rows := libs.Select("*", "places", "id %2 = 0")

	// for i := range rows {
	// 	fmt.Println(rows[i]["address"])
	// }
}