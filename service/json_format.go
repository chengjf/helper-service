package service

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	_ "github.com/pkg/errors"
)

var x = "hello"

type Input struct {
	Input string `json:"input"`
}

func FormatJson(ctx iris.Context) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(r.(error).Error())
		}
	}()

	var input Input
	var err error
	if err = ctx.ReadJSON(&input); err != nil {
		panic(err)
	}
	str := input.Input
	in := []byte(str)
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	out, _ := json.MarshalIndent(raw, "", "    ")
	result := string(out)
	ctx.JSON(iris.Map{
		"message": result,
	})

}
