package main

import (
    "github.com/kataras/iris"
    "github.com/chengjf/helper-service/service"
)

func main() {
    app := iris.Default()
    app.Get("/api/json-format", service.FormatJson)
    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr("127.0.0.1:8080"))
}
