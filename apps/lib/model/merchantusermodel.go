package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MerchantUserModel = (*customMerchantUserModel)(nil)

type (
	// MerchantUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMerchantUserModel.
	MerchantUserModel interface {
		merchantUserModel
	}

	customMerchantUserModel struct {
		*defaultMerchantUserModel
	}
)

// NewMerchantUserModel returns a model for the database table.
func NewMerchantUserModel(conn sqlx.SqlConn, c cache.CacheConf) MerchantUserModel {
	return &customMerchantUserModel{
		defaultMerchantUserModel: newMerchantUserModel(conn, c),
	}
}
