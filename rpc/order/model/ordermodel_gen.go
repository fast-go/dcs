// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderFieldNames          = builder.RawFieldNames(&Order{})
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderRowsWithPlaceHolder = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDcsOrderIdPrefix = "cache:dcs:order:id:"
)

type (
	orderModel interface {
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Order, error)
		Update(ctx context.Context, data *Order) error
		Delete(ctx context.Context, id int64) error

		TxInsert(ctx context.Context, tx *sql.Tx, data *Order) (sql.Result, error)
		TxUpdate(ctx context.Context, tx *sql.Tx, data *Order) error

	}

	defaultOrderModel struct {
		sqlc.CachedConn
		table string
	}

	Order struct {
		Id          int64  `db:"id"`
		ProductName string `db:"product_name"`
		ProductId   int64  `db:"product_id"`
		Uid         int64  `db:"uid"`
		Status      int64  `db:"status"`
		Num         int64  `db:"num"`
	}
)

func newOrderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultOrderModel {
	return &defaultOrderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`order`",
	}
}

func (m *defaultOrderModel) Delete(ctx context.Context, id int64) error {
	dcsOrderIdKey := fmt.Sprintf("%s%v", cacheDcsOrderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, dcsOrderIdKey)
	return err
}

func (m *defaultOrderModel) FindOne(ctx context.Context, id int64) (*Order, error) {
	dcsOrderIdKey := fmt.Sprintf("%s%v", cacheDcsOrderIdPrefix, id)
	var resp Order
	err := m.QueryRowCtx(ctx, &resp, dcsOrderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	dcsOrderIdKey := fmt.Sprintf("%s%v", cacheDcsOrderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductName, data.ProductId, data.Uid, data.Status, data.Num)
	}, dcsOrderIdKey)
	return ret, err
}

func (m *defaultOrderModel) Update(ctx context.Context, data *Order) error {
	dcsOrderIdKey := fmt.Sprintf("%s%v", cacheDcsOrderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProductName, data.ProductId, data.Uid, data.Status, data.Num, data.Id)
	}, dcsOrderIdKey)
	return err
}


func (m *defaultOrderModel) TxInsert(ctx context.Context, tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
	ret, err := tx.ExecContext(ctx, query, data.ProductName, data.ProductId, data.Uid, data.Status, data.Num)

	return ret, err
}

func (m *defaultOrderModel) TxUpdate(ctx context.Context, tx *sql.Tx, data *Order) error {
	productIdKey := fmt.Sprintf("%s%v", cacheDcsOrderIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return tx.ExecContext(ctx, query, data.ProductName, data.ProductId, data.Uid, data.Status, data.Num, data.Id)
	}, productIdKey)
	return err
}


func (m *defaultOrderModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheDcsOrderIdPrefix, primary)
}

func (m *defaultOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderModel) tableName() string {
	return m.table
}

