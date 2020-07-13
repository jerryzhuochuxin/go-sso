package services

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"sso/defs"
	"time"
)

var (
	tokenKeyTemplate = "tokenKey_%s"
	localCache       *cache.Cache
	redisCache       *redis.Client
)

func InitJwtCache() {
	if defs.RedisUrl != "" {
		defs.JwtCache = defs.USE_REDIS_CACHE
		redisCache = redis.NewClient(&redis.Options{
			Addr: defs.RedisUrl,
		})

		if _, err := redisCache.Ping(context.TODO()).Result(); err != nil {
			fmt.Println("redis connect error")
		}

		return
	}

	if defs.Cache {
		defs.JwtCache = defs.USE_LOCAL_CACHE
		localCache = cache.New(5*time.Minute, 10*time.Minute)
		return
	}

	defs.JwtCache = defs.NO_USE_CACHE
}

func ReSetCacheForToken(tokenKey, tokenValue string, expires int) (bool, error) {
	cacheKey := fmt.Sprintf(tokenKeyTemplate, tokenKey)
	if defs.JwtCache == defs.USE_REDIS_CACHE {
		redisCache.Set(context.TODO(), cacheKey, tokenValue, time.Duration(expires)*time.Minute)
	} else {
		localCache.Set(cacheKey, tokenValue, time.Duration(expires)*time.Minute)
	}

	return true, nil
}

func FindKeyFromCache(tokenKey string) string {
	cacheKey := fmt.Sprintf(tokenKeyTemplate, tokenKey)
	if defs.JwtCache == defs.USE_REDIS_CACHE {
		tokenValue, _ := redisCache.Get(context.TODO(), cacheKey).Result()
		return tokenValue
	}

	tokenValue, found := localCache.Get(cacheKey)
	if !found {
		return ""
	}
	return tokenValue.(string)
}

func DeleteKeyFromCache(tokenKey string) {
	cacheKey := fmt.Sprintf(tokenKeyTemplate, tokenKey)
	if defs.JwtCache == defs.USE_REDIS_CACHE {
		redisCache.Del(context.TODO(), cacheKey)
		return
	}
	localCache.Delete(cacheKey)
}
