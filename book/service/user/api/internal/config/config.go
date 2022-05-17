package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// Mysql
	Mysql struct {
		DataSource string
	}
	// Redis
	CacheRedis cache.CacheConf

	// Auth jwt鉴权参数配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
