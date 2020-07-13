package defs

import (
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
	Port       string
	ServiceMap sync.Map
)


