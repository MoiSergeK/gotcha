package controllers

import (
	"github.com/kataras/iris"
	"../libs"
	"../db"
)

func AddFriends(ctx iris.Context) {
	libs.Log("POST:/places/common")

	var req map[string]string
	ctx.ReadJSON(&req)

	db.Insert("common_places", req)

	ctx.JSON(iris.Map{"status": 200, "data": "ok"})
}

func GetFriends(ctx iris.Context) {
	// userId := ctx.Params().Get("id")

	// // db.Insert("location_requests", map[string]string{"userId": userId})

	// ctx.JSON(iris.Map{"status": 200, "data": "ok"})
}

func GetResearchedFriends(ctx iris.Context) {

}

func RequestFriendLocation(ctx iris.Context) {

}

func RequestAllFriendsLocations(ctx iris.Context) {

}

func ShareLocationWithFriend(ctx iris.Context) {

}

func ShareLocationWithAllFriends(ctx iris.Context) {

}

func PushSelfRoute(ctx iris.Context) {
	
}

func GetSelfRoutes(ctx iris.Context) {

}

func GetSelfRoute(ctx iris.Context) {

}

func GetCommonRoutes(ctx iris.Context) {

}

func GetCommonRoute(ctx iris.Context) {
	
}