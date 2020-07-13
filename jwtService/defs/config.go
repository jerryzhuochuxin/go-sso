package defs

var (
	Port             string
	DefaultSecretKey string
	DefaultExpires   int

	Cache    bool
	RedisUrl string

	JwtCache ServerCache
)

type ServerCache int

const (
	NO_USE_CACHE    ServerCache = 0
	USE_LOCAL_CACHE ServerCache = 1
	USE_REDIS_CACHE ServerCache = 2
)

func (s *ServerCache) HasUseCache() bool {
	return *s != NO_USE_CACHE
}
