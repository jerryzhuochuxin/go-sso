package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"jwtService/defs"
	"jwtService/handlers"
	"jwtService/middlewares"
	"jwtService/services"
)

func main() {
	parseParam()
	services.InitJwtCache()

	c := gin.Default()

	c.Use(middlewares.UseJsonResult)
	c.Use(middlewares.UseCross)
	c.Use(middlewares.UseJumpOptionsMethods)

	c.POST("api/getToken", handlers.GetToken)
	c.POST("api/authToken", handlers.AuthToken)
	c.GET("api/ping", handlers.GetPong)

	if defs.JwtCache != defs.NO_USE_CACHE {
		c.POST("api/deleteToken", handlers.DeleteToken)
	}

	c.Run(":" + defs.Port)
}

func parseParam() {
	flag.StringVar(&defs.Port, "port", "8080", "启动端口")
	flag.StringVar(&defs.DefaultSecretKey, "key", "a secret key", "token默认秘钥")
	flag.IntVar(&defs.DefaultExpires, "expires", 5, "token默认过期时间(分)")

	flag.BoolVar(&defs.Cache, "cache", false, "是否缓存jwt的状态")
	flag.StringVar(&defs.RedisUrl, "redisUrl", "", "是否使用redis来缓存jwt状态,如果需要则配置相应的url")

	flag.Parse()
}
