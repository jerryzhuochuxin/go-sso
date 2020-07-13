package defs

import "github.com/dgrijalva/jwt-go"

type JwtInfoClaims struct {
	Data []byte
	jwt.StandardClaims
}

type HttpResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetSuccessHttpResult(data interface{}) HttpResult {
	return HttpResult{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func GetLoginInvalidHttpResult() HttpResult {
	return HttpResult{
		Code:    401,
		Message: "login invalid",
		Data:    nil,
	}
}

func GetBadRequestHttpResult(message string) HttpResult {
	return HttpResult{
		Code:    400,
		Message: message,
		Data:    nil,
	}
}

func GetServerErrorHttpResult(message string) HttpResult {
	return HttpResult{
		Code:    500,
		Message: message,
		Data:    nil,
	}
}
