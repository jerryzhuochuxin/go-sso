package services

import (
	"flag"
	"jwtService/defs"
)

func InitServerConfig() {
	flag.StringVar(&defs.Port, "port", "8080", "启动端口")
	flag.StringVar(&defs.DefaultSecretKey, "key", "a secret key", "token默认秘钥")
	flag.IntVar(&defs.DefaultExpires, "expires", 5, "token默认过期时间(分)")

	flag.BoolVar(&defs.Cache, "cache", false, "是否缓存jwt的状态")
	flag.StringVar(&defs.RedisUrl, "redisUrl", "", "是否使用redis来缓存jwt状态,如果需要则配置相应的url")

	flag.StringVar(&defs.RegisterCenterUrl, "centerUrl", "", "服务注册中心")

	flag.Parse()
}
