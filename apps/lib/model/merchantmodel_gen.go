// Code generated by goctl. DO NOT EDIT!

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
	merchantFieldNames          = builder.RawFieldNames(&Merchant{})
	merchantRows                = strings.Join(merchantFieldNames, ",")
	merchantRowsExpectAutoSet   = strings.Join(stringx.Remove(merchantFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	merchantRowsWithPlaceHolder = strings.Join(stringx.Remove(merchantFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheMerchantIdPrefix = "cache:merchant:id:"
)

type (
	merchantModel interface {
		Insert(ctx context.Context, data *Merchant) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Merchant, error)
		Update(ctx context.Context, newData *Merchant) error
		Delete(ctx context.Context, id int64) error
	}

	defaultMerchantModel struct {
		sqlc.CachedConn
		table string
	}

	Merchant struct {
		Id        int64     `db:"id"`     // merchant id
		Name      string    `db:"name"`   // merchant name
		Status    int64     `db:"status"` // merchant status(0 active,-1 inactive)
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}
	Total struct {
		Total        int64     `db:"total"`
	}
)
func (Merchant) TableName() string {
    return "merchant"
}
func newMerchantModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultMerchantModel {
	return &defaultMerchantModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`merchant`",
	}
}

func (m *defaultMerchantModel) Delete(ctx context.Context, id int64) error {
	merchantIdKey := fmt.Sprintf("%s%v", cacheMerchantIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, merchantIdKey)
	return err
}

func (m *defaultMerchantModel) FindOne(ctx context.Context, id int64) (*Merchant, error) {
	merchantIdKey := fmt.Sprintf("%s%v", cacheMerchantIdPrefix, id)
	var resp Merchant
	err := m.QueryRowCtx(ctx, &resp, merchantIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", merchantRows, m.table)
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

func (m *defaultMerchantModel) Insert(ctx context.Context, data *Merchant) (sql.Result, error) {
	merchantIdKey := fmt.Sprintf("%s%v", cacheMerchantIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, merchantRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Status, data.CreatedAt, data.UpdatedAt)
	}, merchantIdKey)
	return ret, err
}

func (m *defaultMerchantModel) Update(ctx context.Context, data *Merchant) error {
	merchantIdKey := fmt.Sprintf("%s%v", cacheMerchantIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, merchantRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Status, data.CreatedAt, data.UpdatedAt, data.Id)
	}, merchantIdKey)
	return err
}

func (m *defaultMerchantModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMerchantIdPrefix, primary)
}

func (m *defaultMerchantModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", merchantRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultMerchantModel) tableName() string {
	return m.table
}
