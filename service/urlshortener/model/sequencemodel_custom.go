package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func (m *customSequenceModel) Next(ctx context.Context) (uint64, error) {
	query := fmt.Sprintf("replace into %s (stub) values ('a')", m.table)
	fmt.Println(query)

	var stmt sqlx.StmtSession
	stmt, err := m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	ret, err := stmt.ExecCtx(ctx)
	if err != nil {
		return 0, err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println(id)
	return uint64(id), nil
}
