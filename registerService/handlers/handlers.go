package handlers

import (
	"github.com/gin-gonic/gin"
	"registerService/defs"
)

func AddService(c *gin.Context) {
	serviceIp := c.ClientIP() + ":" + c.Param("servicePort")
	serviceType := c.Param("serviceType")
	serviceName := c.Param("serviceName")

	_, err := defs.AddService(serviceIp, serviceType, serviceName)

	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(200, "ok")
}
