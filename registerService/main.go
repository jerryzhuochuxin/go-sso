package main

import (
	"github.com/gin-gonic/gin"
	"registerService/handlers"
	"registerService/services"
)

func main() {
	services.InitSchedule()
	c := gin.Default()

	c.POST("api/addService/:serviceType/:serviceName/:servicePort", handlers.AddService)
	c.GET("api/getServices/:serviceType", handlers.GetServices)
	c.GET("api/getAllServices", handlers.GetAllServices)

	c.Run(":8081")
}
