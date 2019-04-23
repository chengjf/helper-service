package main

import (
    "github.com/kataras/iris"
    "github.com/chengjf/helper-service/service"
)

func main() {
    app := iris.Default()
    app.Get("/json-format", service.FormatJson)
    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":8080"))
}
