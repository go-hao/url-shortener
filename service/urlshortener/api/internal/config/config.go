package config

import (
	"github.com/go-hao/url-shortener/pkg/base62"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	UrlMapCacheRedis cache.CacheConf
	UrlMapSql        struct {
		Dsn string
	}

	SequenceSql struct {
		Dsn string
	}

	Base62 base62.Base62Conf

	ShortUrlBlackList []string
	ShortUrlDomain    string
}
