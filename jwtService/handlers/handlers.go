package handlers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"jwtService/defs"
	"jwtService/services"
	"strconv"
	"strings"
)

func GetToken(c *gin.Context) {
	secretKey := strings.TrimSpace(c.Query("secretKey"))
	expires, _ := strconv.Atoi(c.Query("expires"))

	if secretKey == "" {
		secretKey = defs.DefaultSecretKey
	}
	if expires == 0 {
		expires = defs.DefaultExpires
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	tokenValue := services.GetToken(body, secretKey, expires)

	if defs.JwtCache.HasUseCache() {
		tokenKey := strings.TrimSpace(c.Query("tokenKey"))
		if tokenKey == "" {
			c.JSON(400, defs.GetBadRequestHttpResult("后台服务使用了cache,服务请携带tokenKey"))
			return
		}

		if _, err := services.ReSetCacheForToken(tokenKey, tokenValue, expires); err != nil {
			c.JSON(400, defs.GetServerErrorHttpResult("后台错误请检查后台日志"))
			return
		}
	}

	c.JSON(200, defs.GetSuccessHttpResult(tokenValue))
}

func AuthToken(c *gin.Context) {
	tokenValue := strings.TrimSpace(c.GetHeader("X-TokenValue"))
	if tokenValue == "" {
		c.JSON(400, defs.GetBadRequestHttpResult("未携带X-TokenValue"))
		return
	}

	secretKey := strings.TrimSpace(c.Query("secretKey"))
	if secretKey == "" {
		secretKey = defs.DefaultSecretKey
	}

	if defs.JwtCache.HasUseCache() {
		tokenKey := strings.TrimSpace(c.GetHeader("X-TokenKey"))
		if tokenKey == "" {
			c.JSON(400, defs.GetBadRequestHttpResult("后台服务使用了cache,服务请携带X-TokenKey"))
			return
		}
		currentTokenValue := services.FindKeyFromCache(tokenKey)
		if currentTokenValue == "" || currentTokenValue != tokenValue {
			c.JSON(401, defs.GetLoginInvalidHttpResult())
			return
		}
	}

	data, err := services.AuthToken(tokenValue, secretKey)
	if err != nil {
		c.JSON(401, defs.GetLoginInvalidHttpResult())
		return
	}

	c.String(200, "%s", string(data))
}

func DeleteToken(c *gin.Context) {
	tokenKey := strings.TrimSpace(c.GetHeader("X-TokenKey"))
	if tokenKey == "" {
		c.JSON(400, defs.GetBadRequestHttpResult("后台服务使用了cache,服务请携带tokenKey"))
		return
	}
	services.DeleteKeyFromCache(tokenKey)
	c.JSON(200, defs.GetSuccessHttpResult("删除成功"))
}

func GetPong(c *gin.Context) {
	c.String(200, "pong")
}
