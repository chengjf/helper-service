package main

import "github.com/kataras/iris"

func main() {
    app := iris.Default()
    app.Get("/ping", test)
    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":8080"))
}
