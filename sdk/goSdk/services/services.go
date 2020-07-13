package services

import (
	"encoding/json"
	"github.com/jasonlvhit/gocron"
	"goSdk/defs"
	"io/ioutil"
	"net/http"
)

func InitSchedule() {
	gocron.Every(10).Seconds().Do(CheckRegisterIp)
	go func() {
		<-gocron.Start()
	}()
}

func CheckRegisterIp() {
	response, err := http.Get(defs.RegisterIp + "/api/getAllServices")
	if err != nil {
		panic(err)
	}
	bodyReader := response.Body
	bodyByte, _ := ioutil.ReadAll(bodyReader)
	json.Unmarshal(bodyByte, &defs.ServiceList)
}
