package modelc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UrlMapModel = (*customUrlMapModel)(nil)

type (
	// UrlMapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUrlMapModel.
	UrlMapModel interface {
		urlMapModel
	}

	customUrlMapModel struct {
		*defaultUrlMapModel
	}
)

// NewUrlMapModel returns a model for the database table.
func NewUrlMapModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UrlMapModel {
	return &customUrlMapModel{
		defaultUrlMapModel: newUrlMapModel(conn, c, opts...),
	}
}
