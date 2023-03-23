package main

import (
	"fmt"
	"lark/pkg/common/xjwt"
)

func main() {
	token, err := xjwt.CreateToken(1, 1, false, 3600)
	if err != nil {
		panic(err)
	}
	fmt.Printf("token生成完成，token【%v】", token.Token)
}
