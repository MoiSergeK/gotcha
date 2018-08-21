package app

import (
	"github.com/kataras/iris"
	"./libs"
	"regexp"
	// "encoding/json"
	"os"
)

func Router(app *iris.Application) {
	app.Get("/app.js", func(ctx iris.Context) {
		ctx.WriteString(libs.JS("app.js"))
	})

	app.Get("/", func(ctx iris.Context) {
		libs.Log("/")

		view := libs.View("index")

		for key, value := range map[string]string{ "title": "Gotcha | HOME" } {
			r := regexp.MustCompile("{{ ?" + key + " ?}}")
			view = r.ReplaceAllString(view, value)
		}
		
		ctx.WriteString(view)
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

	app.Post("/places/self", func(ctx iris.Context) {
		libs.Log("POST:/place/self")

		var req map[string]string
		ctx.ReadJSON(&req)

		libs.Insert("places", req)

		ctx.JSON(iris.Map{"status": 200, "data": "ok"})
	})

	app.Post("/places/common", func(ctx iris.Context) {
		libs.Log("POST:/places/common")

		var req map[string]string
		ctx.ReadJSON(&req)

		libs.Println(req["name"])

		libs.Insert("common_places", req)

		ctx.JSON(iris.Map{"status": 200, "data": "ok"})
	})

	app.Get("/places/self", func(ctx iris.Context) {
		libs.Log("GET:/places/self")
		
		ctx.Header("Access-Control-Allow-Origin", "*")

		rows := libs.Select("*", "places", "")

		var data []iris.Map

		for _, value := range rows {
			data = append(data, iris.Map{"id": value["id"], "lat": value["lat"], "lng": value["lng"], "address": value["address"]})
		}

		ctx.JSON(iris.Map{"status": 200, "data": data})
	})

	app.Get("/places/common", func(ctx iris.Context) {
		libs.Log("GET:/places/common")
		
		ctx.Header("Access-Control-Allow-Origin", "*")

		rows := libs.Select("*", "common_places", "")

		var data []iris.Map

		for _, value := range rows {
			data = append(data, iris.Map{"id": value["id"], "lat": value["lat"], "lng": value["lng"], "address": value["address"]})
		}

		ctx.JSON(iris.Map{"status": 200, "data": data})
	})

	app.Options("/*", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		ctx.WriteString("")
	})

	app.Delete("/places/self/{id}", func(ctx iris.Context) {
		libs.Log("DELETE:/places/self/{id}")

		ctx.Header("Access-Control-Allow-Origin", "*")

		libs.Delete("places", "id=" + ctx.Params().Get("id"))

		ctx.JSON(iris.Map{"status": 200, "data": "ok"})
	})

	app.Delete("/places/common/{id}", func(ctx iris.Context) {
		libs.Log("DELETE:/places/common/{id}")

		ctx.Header("Access-Control-Allow-Origin", "*")

		libs.Delete("common_places", "id=" + ctx.Params().Get("id"))

		ctx.JSON(iris.Map{"status": 200, "data": "ok"})
	})

	app.Get("/log", func(ctx iris.Context) {
		data := libs.ReadFileAsText("./app/log/log.txt")

		ctx.JSON(iris.Map{"status": 200, "data": data})
	})

	app.Delete("/log", func(ctx iris.Context) {
		os.Create("./app/log/log.txt")

		ctx.JSON(iris.Map{"status": 200, "data": "ok"})
	})
}