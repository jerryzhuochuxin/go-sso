package main

import (
	"github.com/gin-gonic/gin"
	"jwtService/defs"
	"jwtService/handlers"
	"jwtService/middlewares"
	"jwtService/services"
)

func main() {
	services.InitServerConfig()
	services.InitJwtCache()
	services.InitRegisterCenter()

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
