package service

import (
	"github.com/kataras/iris"
	"encoding/json"
	_ "fmt"
)

var x = "hello"

func FormatJson(ctx iris.Context) {

	str := `{
		"str": "foo",
		"num": 100,
		"bool": false,
		"null": null,
		"array": ["foo", "bar", "baz"],
		"obj": { "a": 1, "b": 2 }
	  }`
	  in := []byte(str)
	  var raw map[string]interface{}
	  json.Unmarshal(in, &raw)
	  out, _ := json.MarshalIndent(raw,"","    ")
	  result := string(out);
	  ctx.JSON(iris.Map{
		"message": result,
	})
}