package middlewares

import (
	"github.com/gin-gonic/gin"
	"jwtService/defs"
)

func UseJsonResult(c *gin.Context) {
	c.Header("content-type", "application/json")
	c.Next()
}

func UseCross(c *gin.Context) {
	c.Header("access-control-allow-origin", "*")
	c.Header("access-control-allow-headers", "*")
	c.Header("access-control-allow-methods", "*")
	c.Next()
}
func UseJumpOptionsMethods(c *gin.Context) {
	method := c.Request.Method
	if method == "OPTIONS" {
		c.JSON(200, defs.GetSuccessHttpResult("options method"))
		return
	}

	c.Next()
}
