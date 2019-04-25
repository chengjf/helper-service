package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/kataras/iris"
	"github.com/pkg/errors"
)

type Md5Input struct {
	Input string `json:"input"`
	Count int    `json:"count"`
}

func Md5(ctx iris.Context) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(r.(error).Error())
		}
	}()

	var input Md5Input
	var err error
	if err = ctx.ReadJSON(&input); err != nil {
		panic(err)
	}

	str := input.Input
	count := input.Count
	if count <= 0 || count >= 10 {
		panic(errors.New("Count must between 0 and 10"))
	}

	data := md5.Sum([]byte(str))

	for i := 1; i < count; i++ {
		x := []byte(hex.EncodeToString(data[:]))
		data = md5.Sum(x)
		fmt.Printf("data: %+v\n", data)

	}
	fmt.Printf("last data: %+v\n", data)
	ctx.JSON(iris.Map{
		"result": result,
	})

}
