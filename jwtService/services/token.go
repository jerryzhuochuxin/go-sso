package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"jwtService/defs"
	"time"
)

func GetToken(data []byte, key string, expires int) string {
	claims := defs.JwtInfoClaims{Data: data}
	claims.ExpiresAt = time.Now().Add(time.Duration(expires) * time.Minute).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString([]byte(key))

	if err != nil {
		logrus.Warn(err.Error())
	}

	return token
}

func AuthToken(token string, secretKey string) ([]byte, error) {
	tokenJwt, err := jwt.ParseWithClaims(token, &defs.JwtInfoClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !tokenJwt.Valid {
		return nil, errors.New("login invalid")
	}
	data := tokenJwt.Claims.(*defs.JwtInfoClaims).Data
	return data, nil
}
