package service

import (
	"github.com/kataras/iris"
	"encoding/json"
	_ "fmt"
)

var x = "hello"

type Input struct {
    Input      string     `json:"input"`
}

func FormatJson(ctx iris.Context) {

	// str := `{
	// 	"str": "foo",
	// 	"num": 100,
	// 	"bool": false,
	// 	"null": null,
	// 	"array": ["foo", "bar", "baz"],
	// 	"obj": { "a": 1, "b": 2 }
	//   }`
	var input Input
	if err := ctx.ReadJSON(&input); err != nil {
		// Handle error.
	}
	str := input.Input
	in := []byte(str)
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	out, _ := json.MarshalIndent(raw,"","    ")
	result := string(out);
	ctx.JSON(iris.Map{
	"message": result,
})
}