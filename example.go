package main

import (
	"github.com/chengjf/helper-service/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pprof"
)

func main() {
	app := iris.Default()
    app.Post("/api/json-format", service.FormatJson)
	app.Post("/api/md5", service.Md5)
	
	// add pprof
	app.Any("/api/debug/pprof/{action:path}", pprof.New())

	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr("127.0.0.1:8080"))
}
