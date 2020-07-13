package goSdk

import (
	"errors"
	"github.com/sirupsen/logrus"
	"goSdk/defs"
	"goSdk/services"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

func SetRegisterIp(ip string) {
	defs.RegisterIp = ip
	services.CheckRegisterIp()
	services.InitSchedule()
}

func GetIpByServiceType(serviceType string) (string, error) {
	if defs.RegisterIp == "" || defs.ServiceList == nil {
		panic("没有设置registerIp")
	}

	ipList, found := defs.ServiceList[serviceType]
	if !found {
		logrus.Warn(serviceType + " 所对应的服务不存在")
		return "", errors.New(serviceType + " 所对应的服务不存在")
	}

	ipIndex := rand.Intn(len(ipList))
	serviceIp := ipList[ipIndex]
	return serviceIp, nil
}

func GetResponse(method, url string, body io.Reader, options func(req *http.Request)) (*http.Response, error) {
	serviceTypeEndIndex := strings.Index(url, "/")

	serviceIp, err := GetIpByServiceType(url[:serviceTypeEndIndex])
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, "http://"+serviceIp+url[serviceTypeEndIndex:], body)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}

	if options != nil {
		options(req)
	}

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}

	return res, nil
}
