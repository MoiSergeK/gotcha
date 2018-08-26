package controllers

import (
	"github.com/kataras/iris"
	"../libs"
	"../db"
	"fmt"
)

func AddSelfPlaces(ctx iris.Context) {
	libs.Log("POST:/places/self")

	var req map[string]string
	ctx.ReadJSON(&req)

	db.Insert("places", req)

	ctx.JSON(iris.Map{"status": 200, "data": "ok"})
}

func AddCommonPlaces(ctx iris.Context) {
	libs.Log("POST:/places/common")

	var req map[string]string
	ctx.ReadJSON(&req)

	db.Insert("common_places", req)

	ctx.JSON(iris.Map{"status": 200, "data": "ok"})
}

func GetSelfPlaces(ctx iris.Context) {
	libs.Log("GET:/places/self")
		
	ctx.Header("Access-Control-Allow-Origin", "*")

	rows := db.Select("*", "places", "")

	var data []iris.Map

	for _, value := range rows {
		data = append(data, iris.Map{"id": value["id"], "lat": value["lat"], "lng": value["lng"], "address": value["address"]})
	}

	ctx.JSON(iris.Map{"status": 200, "data": data})
}

func GetCommonPlaces(ctx iris.Context) {
	libs.Log("GET:/places/common")
		
	ctx.Header("Access-Control-Allow-Origin", "*")

	rows := db.Select("*", "common_places", "")

	var data []iris.Map

	for _, value := range rows {
		data = append(data, iris.Map{"id": value["id"], "lat": value["lat"], "lng": value["lng"], "name": value["name"]})
	}

	ctx.JSON(iris.Map{"status": 200, "data": data})
}

func DeleteSelfPlace(ctx iris.Context) {
	libs.Log("DELETE:/places/self/{id}")

	ctx.Header("Access-Control-Allow-Origin", "*")

	db.Delete("places", "id=" + ctx.Params().Get("id"))

	ctx.JSON(iris.Map{"status": 200, "data": "ok"})
}

func DeleteCommonPlace(ctx iris.Context) {
	libs.Log("DELETE:/places/common/{id}")

	ctx.Header("Access-Control-Allow-Origin", "*")

	db.Delete("common_places", "id=" + ctx.Params().Get("id"))

	ctx.JSON(iris.Map{"status": 200, "data": "ok"})
}

func GetNearbyPlaces(ctx iris.Context) {
	lattitude,_ := ctx.URLParamFloat64("lat")
	longitude,_ := ctx.URLParamFloat64("lng")

	lat := fmt.Sprintf("%f", lattitude)
	lng := fmt.Sprintf("%f", longitude)

	rows := db.Select("*", "common_places", "abs(lat-" + string(lat) + ") <= 0.001 and abs(lng-" + string(lng) + ") <= 0.001")

	ctx.JSON(iris.Map{"status": 200, "data": rows})
}