package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lark/pkg/common/xjwt"
	"lark/pkg/entity/vo/response"
)

func main() {
	engine := gin.Default()
	engine.Use(JwtAuth(), TestHandlerFunc())
	engine.GET("hello", hello)
	engine.Run(":8080")
}

func hello(ctx *gin.Context) {
	fmt.Println("访问了hello的api")
	ctx.SecureJSON(200, response.Ok)
}

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := xjwt.ParseFromCookie(ctx)
		if err != nil {
			ctx.Abort()
			ctx.SecureJSON(-1, response.AppErr.WithMsg("请登录"))
			return
		}
		fmt.Println("验证成功", token)
	}
}

func TestHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("测试中间件")
	}
}
