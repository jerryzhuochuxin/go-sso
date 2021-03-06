package handlers

import (
	"github.com/gin-gonic/gin"
	"registerService/services"
)

func AddService(c *gin.Context) {
	serviceIp := c.ClientIP() + ":" + c.Param("servicePort")
	serviceType := c.Param("serviceType")
	serviceName := c.Param("serviceName")

	_, err := services.AddService(serviceIp, serviceType, serviceName)

	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(200, "ok")
}

func GetServices(c *gin.Context) {
	serviceType := c.Param("serviceType")
	ipList := services.SelectServicesByType(serviceType)
	c.JSON(200, ipList)
}

func GetAllServices(c *gin.Context) {
	c.JSON(200, services.SelectAllServices())
}
