package svc

import (
	"github.com/go-hao/url-shortener/pkg/base62"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/config"
	"github.com/go-hao/url-shortener/service/urlshortener/model"
	"github.com/go-hao/url-shortener/service/urlshortener/modelc"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ErrBadReqest any

	UrlMapModel   modelc.UrlMapModel
	SequenceModel model.SequenceModel

	ShortUrlBlackList map[string]struct{}

	Base62 *base62.Base62
	Filter *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	urlMapConn := sqlx.NewMysql(c.UrlMapSql.Dsn)
	sequenceConn := sqlx.NewMysql(c.SequenceSql.Dsn)

	store := redis.MustNewRedis(c.BloomRedis)

	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}
	base62 := base62.MustNew(c.Base62)

	return &ServiceContext{
		Config:       c,
		ErrBadReqest: nil,

		UrlMapModel:   modelc.NewUrlMapModel(urlMapConn, c.UrlMapCacheRedis),
		SequenceModel: model.NewSequenceModel(sequenceConn),

		ShortUrlBlackList: m,

		Base62: base62,
		Filter: bloom.New(store, "bloom", 20*(1<<20)),
	}
}
