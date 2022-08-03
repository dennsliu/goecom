// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	merchantUserFieldNames          = builder.RawFieldNames(&MerchantUser{})
	merchantUserRows                = strings.Join(merchantUserFieldNames, ",")
	merchantUserRowsExpectAutoSet   = strings.Join(stringx.Remove(merchantUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	merchantUserRowsWithPlaceHolder = strings.Join(stringx.Remove(merchantUserFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheMerchantUserIdPrefix       = "cache:merchantUser:id:"
	cacheMerchantUserEmailPrefix    = "cache:merchantUser:email:"
	cacheMerchantUserUsernamePrefix = "cache:merchantUser:username:"
)

type (
	merchantUserModel interface {
		Insert(ctx context.Context, data *MerchantUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*MerchantUser, error)
		FindOneByEmail(ctx context.Context, email string) (*MerchantUser, error)
		FindOneByUsername(ctx context.Context, username string) (*MerchantUser, error)
		Update(ctx context.Context, newData *MerchantUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultMerchantUserModel struct {
		sqlc.CachedConn
		table string
	}

	MerchantUser struct {
		Id          int64          `db:"id"` // merchant user id
		Nickname    sql.NullString `db:"nickname"`
		Username    string         `db:"username"`
		Email       string         `db:"email"`
		Password    string         `db:"password"`
		Telephone   sql.NullString `db:"telephone"`
		Mobliephone sql.NullString `db:"mobliephone"`
		MerchantId  int64          `db:"merchant_id"` // merchantuser
		Status      int64          `db:"status"`
		CreatedAt   time.Time      `db:"created_at"`
		UpdatedAt   time.Time      `db:"updated_at"`
	}
)

func newMerchantUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultMerchantUserModel {
	return &defaultMerchantUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`merchant_user`",
	}
}

func (m *defaultMerchantUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	merchantUserEmailKey := fmt.Sprintf("%s%v", cacheMerchantUserEmailPrefix, data.Email)
	merchantUserIdKey := fmt.Sprintf("%s%v", cacheMerchantUserIdPrefix, id)
	merchantUserUsernameKey := fmt.Sprintf("%s%v", cacheMerchantUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, merchantUserEmailKey, merchantUserIdKey, merchantUserUsernameKey)
	return err
}

func (m *defaultMerchantUserModel) FindOne(ctx context.Context, id int64) (*MerchantUser, error) {
	merchantUserIdKey := fmt.Sprintf("%s%v", cacheMerchantUserIdPrefix, id)
	var resp MerchantUser
	err := m.QueryRowCtx(ctx, &resp, merchantUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", merchantUserRows, m.table)
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

func (m *defaultMerchantUserModel) FindOneByEmail(ctx context.Context, email string) (*MerchantUser, error) {
	merchantUserEmailKey := fmt.Sprintf("%s%v", cacheMerchantUserEmailPrefix, email)
	var resp MerchantUser
	err := m.QueryRowIndexCtx(ctx, &resp, merchantUserEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", merchantUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMerchantUserModel) FindOneByUsername(ctx context.Context, username string) (*MerchantUser, error) {
	merchantUserUsernameKey := fmt.Sprintf("%s%v", cacheMerchantUserUsernamePrefix, username)
	var resp MerchantUser
	err := m.QueryRowIndexCtx(ctx, &resp, merchantUserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", merchantUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMerchantUserModel) Insert(ctx context.Context, data *MerchantUser) (sql.Result, error) {
	merchantUserEmailKey := fmt.Sprintf("%s%v", cacheMerchantUserEmailPrefix, data.Email)
	merchantUserIdKey := fmt.Sprintf("%s%v", cacheMerchantUserIdPrefix, data.Id)
	merchantUserUsernameKey := fmt.Sprintf("%s%v", cacheMerchantUserUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, merchantUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Nickname, data.Username, data.Email, data.Password, data.Telephone, data.Mobliephone, data.MerchantId, data.Status, data.CreatedAt, data.UpdatedAt)
	}, merchantUserEmailKey, merchantUserIdKey, merchantUserUsernameKey)
	return ret, err
}

func (m *defaultMerchantUserModel) Update(ctx context.Context, newData *MerchantUser) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	merchantUserEmailKey := fmt.Sprintf("%s%v", cacheMerchantUserEmailPrefix, data.Email)
	merchantUserIdKey := fmt.Sprintf("%s%v", cacheMerchantUserIdPrefix, data.Id)
	merchantUserUsernameKey := fmt.Sprintf("%s%v", cacheMerchantUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, merchantUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Nickname, newData.Username, newData.Email, newData.Password, newData.Telephone, newData.Mobliephone, newData.MerchantId, newData.Status, newData.CreatedAt, newData.UpdatedAt, newData.Id)
	}, merchantUserEmailKey, merchantUserIdKey, merchantUserUsernameKey)
	return err
}

func (m *defaultMerchantUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMerchantUserIdPrefix, primary)
}

func (m *defaultMerchantUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", merchantUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultMerchantUserModel) tableName() string {
	return m.table
}
