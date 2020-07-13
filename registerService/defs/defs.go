package defs

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
)

type ServiceObj struct {
	Name        string        `json:"name"`
	Ip          string        `json:"ip"`
	ServiceType string        `json:"serviceType"`
	Status      ServiceStatus `json:"-"`
}

type ServiceStatus int

const (
	ALIVE   ServiceStatus = 0
	MISSING ServiceStatus = 1
	DEAD    ServiceStatus = 2
)

var (
	ServiceMap sync.Map
)

func AddService(ip, serviceType, name string) (bool, error) {
	newService := ServiceObj{
		Name:        name,
		Ip:          ip,
		ServiceType: serviceType,
		Status:      ALIVE,
	}

	_, found := ServiceMap.LoadOrStore(name, newService)
	if !found {
		fmt.Printf("成功注册一个服务:%#v\n", newService)
		return true, nil
	}
	return false, errors.New("已经存在相同的名字服务")
}
func CheckService() {
	ServiceMap.Range(func(k, v interface{}) bool {
		service := v.(ServiceObj)
		res, err := http.Get("http://" + service.Ip + "/api/ping")

		if err != nil || res.StatusCode != 200 {
			ServiceMap.Delete(k.(string))
			fmt.Printf("删除服务%#v\n", service)
		} else {
			fmt.Printf("服务存活%#v\n", service)
		}
		return true
	})
}

func SelectServicesByType(serviceType string) []string {
	var result []string
	ServiceMap.Range(func(k, v interface{}) bool {
		service := v.(ServiceObj)
		if service.ServiceType == serviceType {
			result = append(result, service.Ip)
		}
		return true
	})
	return result
}

func SelectAllServices() map[string][]string {
	result := map[string][]string{}
	ServiceMap.Range(func(k, v interface{}) bool {
		service := v.(ServiceObj)
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
