package services

import (
	"flag"
	"registerService/defs"
)

func InitConfig() {
	flag.StringVar(&defs.Port, "port", "9081", "启动端口号")
	flag.Parse()
}
