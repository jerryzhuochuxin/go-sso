package services

import "github.com/jasonlvhit/gocron"

func InitSchedule() {
	gocron.Every(10).Seconds().Do(CheckService)
	go func() {
		<-gocron.Start()
	}()
}
