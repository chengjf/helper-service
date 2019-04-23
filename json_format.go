package main

import "github.com/kataras/iris"

func test(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "pong",
	})
}