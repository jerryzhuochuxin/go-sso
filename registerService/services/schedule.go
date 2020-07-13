package services

import (
	"github.com/jasonlvhit/gocron"
	"registerService/defs"
)

func InitSchedule() {
	gocron.Every(10).Seconds().Do(defs.CheckService)
	go func() {
		<-gocron.Start()
	}()
}
