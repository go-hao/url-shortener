package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UrlMapModel = (*customUrlMapModel)(nil)

type (
	// UrlMapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUrlMapModel.
	UrlMapModel interface {
		urlMapModel
		withSession(session sqlx.Session) UrlMapModel
	}

	customUrlMapModel struct {
		*defaultUrlMapModel
	}
)

// NewUrlMapModel returns a model for the database table.
func NewUrlMapModel(conn sqlx.SqlConn) UrlMapModel {
	return &customUrlMapModel{
		defaultUrlMapModel: newUrlMapModel(conn),
	}
}

func (m *customUrlMapModel) withSession(session sqlx.Session) UrlMapModel {
	return NewUrlMapModel(sqlx.NewSqlConnFromSession(session))
}
