package handlers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"sso/defs"
	"sso/services"
	"strconv"
	"strings"
)

var (
	DefaultSecretKey string
	DefaultExpires   int
)

func GetToken(c *gin.Context) {
	secretKey := strings.TrimSpace(c.Query("secretKey"))
	expires, _ := strconv.Atoi(c.Query("expires"))

	if secretKey == "" {
		secretKey = DefaultSecretKey
	}
	if expires == 0 {
		expires = DefaultExpires
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.JSON(200, defs.GetSuccessHttpResult(services.GetToken(body, secretKey, expires)))
}

func AuthToken(c *gin.Context) {
	reqToken := strings.TrimSpace(c.GetHeader("X-Token"))
	if reqToken == "" {
		c.JSON(401, defs.GetLoginInvalidHttpResult())
		return
	}

	secretKey := strings.TrimSpace(c.Query("secretKey"))
	if secretKey == "" {
		secretKey = DefaultSecretKey
	}

	data, err := services.AuthToken(reqToken, secretKey)
	if err != nil {
		c.JSON(401, defs.GetLoginInvalidHttpResult())
		return
	}

	c.String(200, "%s", string(data))
}
