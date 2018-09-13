package app

import (
	"github.com/kataras/iris"
	"./libs"
	"./db"
	"os"
	"./controllers"
)

func Router(app *iris.Application) {
	app.Get("/app.js", func(ctx iris.Context) {
		controllers.JS(ctx)
	})

	app.Get("/", func(ctx iris.Context) {
		controllers.Index(ctx)
	})

	app.Options("/*", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		ctx.WriteString("")
	})

	// app.Get("/geocoder", func(ctx iris.Context) {
	// 	lat := ctx.URLParam("lat")
	// 	lng := ctx.URLParam("lng")

	// 	libs.Log("/geocoder?lat=" + lat + "&lng=" + lng)

	// 	// req := libs.GET("https://api.opencagedata.com/geocode/v1/json?q=" + lat + "+" + lng + "&key=efb930759328444aa76ac5123639248c&language=ru&pretty=1")

	// 	// var address string

	// 	// if len(req["results"].([]interface{})) > 0 {
	// 	// 	address = req["results"].([]interface{})[0].(map[string]interface{})["formatted"].(string)
	// 	// } else {
	// 	// 	address = "undefined"
	// 	// }

	// 	geocoder := opencagedata.NewGeocoder("efb930759328444aa76ac5123639248c")
	// 	result, err := geocoder.Geocode(lat + "," + lng, nil)

	// 	var address string

	// 	if err == nil {
	// 		f_result := result.Results[0]
	// 		address = f_result.Formatted
	// 	} else {
	// 		address = "undefined"
	// 	}

	// 	ctx.JSON(iris.Map{"address": address})
	// })

	app.Post("/signup", func(ctx iris.Context) {
		controllers.SignUp(ctx)
	})

	app.Get("/users", func(ctx iris.Context) {
		controllers.GetAllUsers(ctx)
	})

	app.Get("/users/search", func(ctx iris.Context) {
		controllers.SearchUsers(ctx)
	})

	app.Post("/users", func(ctx iris.Context) {
		controllers.AddUser(ctx)
	})

	app.Post("/places/self", func(ctx iris.Context) {
		controllers.AddSelfPlaces(ctx)
	})

	app.Post("/places/common", func(ctx iris.Context) {
		controllers.AddCommonPlaces(ctx)
	})

	app.Get("/places/self", func(ctx iris.Context) {
		controllers.GetSelfPlaces(ctx)
	})

	app.Get("/places/common", func(ctx iris.Context) {
		controllers.GetCommonPlaces(ctx)
	})

	app.Delete("/places/self/{id}", func(ctx iris.Context) {
		controllers.DeleteSelfPlace(ctx)
	})

	app.Delete("/places/common/{id}", func(ctx iris.Context) {
		controllers.DeleteCommonPlace(ctx)
	})

	app.Get("/log", func(ctx iris.Context) {
		data := libs.ReadFileAsText("./app/log/log.txt")
		ctx.JSON(iris.Map{"status": 200, "data": data})
	})

	app.Delete("/log", func(ctx iris.Context) {
		os.Create("./app/log/log.txt")

		ctx.JSON(iris.Map{"status": 200, "data": "ok"})
	})

	app.Post("/app/terminate", func(ctx iris.Context) {
		db.Update("user", "id=1", "online=0")
	})

	app.Get("places/nearby", func(ctx iris.Context) {
		controllers.GetNearbyPlaces(ctx)
	})

	app.Post("friends", func(ctx iris.Context) {
		controllers.AddFriends(ctx)
	})

	app.Get("friends", func(ctx iris.Context) {
		controllers.GetFriends(ctx)
	})

	app.Get("friends/search", func(ctx iris.Context) {
		controllers.GetResearchedFriends(ctx)
	})

	// Request friend's location
	app.Get("friends/{id}/location", func(ctx iris.Context) {
		controllers.RequestFriendLocation(ctx)
	})

	// Share my location with friend
	app.Post("friends/{id}/location", func(ctx iris.Context) {
		controllers.ShareLocationWithFriend(ctx)
	})

	// Request all friends locations
	app.Get("friends/location", func(ctx iris.Context) {
		controllers.RequestAllFriendsLocations(ctx)
	})

	// Share my location with all friends
	app.Post("friends/location", func(ctx iris.Context) {
		controllers.ShareLocationWithAllFriends(ctx)
	})

	// Update self route
	app.Put("route/self", func(ctx iris.Context) {
		controllers.PushSelfRoute(ctx)
	})

	// Get self routes
	app.Get("routes/self", func(ctx iris.Context) {
		controllers.GetSelfRoutes(ctx)
	})

	// Get self route
	app.Get("routes/self/{id}", func(ctx iris.Context) {
		controllers.GetSelfRoute(ctx)
	})

	// Get common routes
	app.Get("routes/common", func(ctx iris.Context) {
		controllers.GetCommonRoutes(ctx)
	})

	// Get common route
	app.Get("routes/common/{id}", func(ctx iris.Context) {
		controllers.GetCommonRoute(ctx)
	})
}