package controllers

import (
	"github.com/kataras/iris"
	"../libs"
	"regexp"
	"../db"
	"strconv"
)


func JS(ctx iris.Context) {
	ctx.WriteString(libs.JS("app.js"))
}

func Index(ctx iris.Context) {
	libs.Log("/")

	view := libs.View("index")

	for key, value := range map[string]string{ "title": "Gotcha | HOME" } {
		r := regexp.MustCompile("{{ ?" + key + " ?}}")
		view = r.ReplaceAllString(view, value)
	}
	
	ctx.WriteString(view)
}

func GetAllUsers(ctx iris.Context) {
	libs.Log("GET:/users")
		
	ctx.Header("Access-Control-Allow-Origin", "*")

	rows := db.Select("id, name, phone", "users", "")

	var data []iris.Map

	for _, value := range rows {
		data = append(data, iris.Map{"id": value["id"], "name": value["name"], "phone": value["phone"]})
	}

	ctx.JSON(iris.Map{"status": 200, "data": data})
}

func SearchUsers(ctx iris.Context) {
	libs.Log("GET:/users/search")
		
	ctx.Header("Access-Control-Allow-Origin", "*")

	tel, _ := ctx.URLParamInt("tel")

	rows := db.Select("id, name, phone", "users", "phone = '+' + " +  strconv.Itoa(tel))

	var data []iris.Map

	for _, value := range rows {
		data = append(data, iris.Map{"id": value["id"], "name": value["name"], "phone": value["phone"]})
	}

	ctx.JSON(iris.Map{"status": 200, "data": data})
}

func AddUser(ctx iris.Context) {
	libs.Log("POST:/users")

	var req map[string]string
	ctx.ReadJSON(&req)

	db.Insert("users", req)

	ctx.JSON(iris.Map{"status": 200, "data": "ok"})
}