// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productFieldNames          = builder.RawFieldNames(&Product{})
	productRows                = strings.Join(productFieldNames, ",")
	productRowsExpectAutoSet   = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	productRowsWithPlaceHolder = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDcsProductIdPrefix = "cache:dcs:product:id:"
)

type (
	productModel interface {
		RowBuilder() squirrel.SelectBuilder
		Insert(ctx context.Context, data *Product) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Product, error)
		Update(ctx context.Context, data *Product) error
		Delete(ctx context.Context, id int64) error

		TxAdjustStock(ctx context.Context, tx *sql.Tx, id int64, delta int) (sql.Result, error)
	}

	defaultProductModel struct {
		sqlc.CachedConn
		table string
	}

	Product struct {
		Id    int64  `db:"id"`
		Name  string `db:"name"`
		Price int64  `db:"price"`
		Stock int64  `db:"stock"`
	}
)

func newProductModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductModel {
	return &defaultProductModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`product`",
	}
}

func (m *defaultProductModel) Delete(ctx context.Context, id int64) error {
	dcsProductIdKey := fmt.Sprintf("%s%v", cacheDcsProductIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, dcsProductIdKey)
	return err
}

func (m *defaultProductModel) FindOne(ctx context.Context, id int64) (*Product, error) {
	dcsProductIdKey := fmt.Sprintf("%s%v", cacheDcsProductIdPrefix, id)
	var resp Product
	err := m.QueryRowCtx(ctx, &resp, dcsProductIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query, _, err := m.RowBuilder().Where(squirrel.Eq{"id": id}).ToSql()
		if err != nil {
			return err
		}
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductModel) Insert(ctx context.Context, data *Product) (sql.Result, error) {
	dcsProductIdKey := fmt.Sprintf("%s%v", cacheDcsProductIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, productRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Price, data.Stock)
	}, dcsProductIdKey)
	return ret, err
}

func (m *defaultProductModel) Update(ctx context.Context, data *Product) error {
	dcsProductIdKey := fmt.Sprintf("%s%v", cacheDcsProductIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Price, data.Stock, data.Id)
	}, dcsProductIdKey)
	return err
}

func (m *defaultProductModel) TxAdjustStock(ctx context.Context, tx *sql.Tx, id int64, delta int) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheDcsProductIdPrefix, id)
	return m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set stock=stock+? where id=? and stock >= -?", m.table)
		return tx.ExecContext(ctx, query, delta, id, delta)
	}, productIdKey)
}

func (m *defaultProductModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheDcsProductIdPrefix, primary)
}

func (m *defaultProductModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

// export logic
func (m *defaultProductModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(productRows).From(m.table)
}


func (m *defaultProductModel) tableName() string {
	return m.table
}
