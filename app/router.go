package app

import (
	"github.com/kataras/iris"
	"./libs"
	"github.com/rubenv/opencagedata"
)

func Router(app *iris.Application) {
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/geocoder", func(ctx iris.Context) {
		lat := ctx.URLParam("lat")
		lng := ctx.URLParam("lng")

		libs.Log("/geocoder?lat=" + lat + "&lng=" + lng)

		// req := libs.GET("https://api.opencagedata.com/geocode/v1/json?q=" + lat + "+" + lng + "&key=efb930759328444aa76ac5123639248c&language=ru&pretty=1")

		// var address string

		// if len(req["results"].([]interface{})) > 0 {
		// 	address = req["results"].([]interface{})[0].(map[string]interface{})["formatted"].(string)
		// } else {
		// 	address = "undefined"
		// }

		geocoder := opencagedata.NewGeocoder("efb930759328444aa76ac5123639248c")
		result, err := geocoder.Geocode(lat + "," + lng, nil)

		var address string

		if err == nil {
			f_result := result.Results[0]
			address = f_result.Formatted
		} else {
			address = "undefined"
		}

		ctx.JSON(iris.Map{"address": address})
	})

	app.Post("/me", func(ctx iris.Context) {
		libs.Log("/me")

		var req map[string]interface{}
		ctx.ReadJSON(&req)

		ctx.JSON(iris.Map{"data": req["data"]})
	})
}