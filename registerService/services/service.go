package services

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"registerService/defs"
)

func AddService(ip, serviceType, name string) (bool, error) {
	newService := defs.ServiceObj{
		Name:        name,
		Ip:          ip,
		ServiceType: serviceType,
		Status:      defs.ALIVE,
	}

	_, found := defs.ServiceMap.LoadOrStore(name, newService)
	if !found {
		logrus.Info(fmt.Sprintf("成功注册一个服务:%#v\n", newService))
		return true, nil
	}

	errMessage := "已经存在相同的名字服务"
	logrus.Warn(errMessage)
	return false, errors.New(errMessage)
}
func CheckService() {
	defs.ServiceMap.Range(func(k, v interface{}) bool {
		service := v.(defs.ServiceObj)
		res, err := http.Get("http://" + service.Ip + "/api/ping")

		if err != nil || res.StatusCode != 200 {
			defs.ServiceMap.Delete(k.(string))
			logrus.Info(fmt.Sprintf("删除服务:%#v\n", service))
		} else {
			logrus.Info(fmt.Sprintf("服务存活:%#v\n", service))
		}
		return true
	})
}

func SelectServicesByType(serviceType string) []string {
	var result []string
	defs.ServiceMap.Range(func(k, v interface{}) bool {
		service := v.(defs.ServiceObj)
		if service.ServiceType == serviceType {
			result = append(result, service.Ip)
		}
		return true
	})
	return result
}

func SelectAllServices() map[string][]string {
	result := map[string][]string{}
	defs.ServiceMap.Range(func(k, v interface{}) bool {
		service := v.(defs.ServiceObj)
		serviceType := service.ServiceType

		_, found := result[serviceType]
		if !found {
			result[serviceType] = []string{}
		}

		result[serviceType] = append(result[serviceType], service.Ip)

		return true
	})
	return result
}
