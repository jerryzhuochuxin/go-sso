package services

import (
	"io/ioutil"
	"jwtService/defs"
	"net/http"
)

func InitRegisterCenter() {
	if defs.RegisterCenterUrl != "" {
		res, err := http.Post(defs.RegisterCenterUrl+"/"+defs.Port, "application/json", nil)
		if err != nil {
			panic("register center connect error")
		}
		if res.StatusCode != 200 {
			reason, _ := ioutil.ReadAll(res.Body)
			panic(string(reason))
		}
	}
}
