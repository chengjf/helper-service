package service

import (
	"bytes"
	"encoding/json"
	_ "fmt"
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
			// fmt.Printf("Panic: %+v\n", r)
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(r.(error).Error())
		}
	}()

	var input Input
	var err error
	if err = ctx.ReadJSON(&input); err != nil {
		panic(err)
	}

	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, []byte(input.Input), "", "    "); err != nil {
		panic(err)
	}

	result := prettyJSON.String()
	// fmt.Printf(result)
	ctx.JSON(iris.Map{
		"message": result,
	})

}
