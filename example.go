package main

import (
	"github.com/chengjf/helper-service/service"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Post("/api/json-format", service.FormatJson)
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr("127.0.0.1:8080"))
}
