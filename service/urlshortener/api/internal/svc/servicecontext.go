package svc

import (
	"github.com/go-hao/url-shortener/pkg/base62"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/config"
	"github.com/go-hao/url-shortener/service/urlshortener/model"
	"github.com/go-hao/url-shortener/service/urlshortener/modelc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UrlMapModel   modelc.UrlMapModel
	SequenceModel model.SequenceModel

	ShortUrlBlackList map[string]struct{}

	Base62 *base62.Base62
}

func NewServiceContext(c config.Config) *ServiceContext {
	urlMapConn := sqlx.NewMysql(c.UrlMapSql.Dsn)
	sequenceConn := sqlx.NewMysql(c.SequenceSql.Dsn)

	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}
	base62 := base62.MustNew(c.Base62)

	return &ServiceContext{
		Config: c,

		UrlMapModel:   modelc.NewUrlMapModel(urlMapConn, c.UrlMapCacheRedis),
		SequenceModel: model.NewSequenceModel(sequenceConn),

		ShortUrlBlackList: m,

		Base62: base62,
	}
}
