package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"sso/handlers"
	"sso/middlewares"
)

var (
	Port string
)

func main() {
	parseParam()
	c := gin.Default()

	c.Use(middlewares.UseJsonResult)
	c.Use(middlewares.UseCorss)
	c.Use(middlewares.UseJumpOptionsMethods)

	c.POST("api/getToken", handlers.GetToken)
	c.POST("api/authToken", handlers.AuthToken)

	c.Run(":" + Port)
}

func parseParam() {
	flag.StringVar(&Port, "port", "8080", "启动端口")
	flag.StringVar(&handlers.DefaultSecretKey, "key", "a secret key jerryzhuoKey", "token默认秘钥")
	flag.IntVar(&handlers.DefaultExpires, "expires", 5, "token默认过期时间(分)")
	flag.Parse()
}
