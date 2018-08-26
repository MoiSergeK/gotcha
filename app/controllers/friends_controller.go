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