package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductBrandModel = (*customProductBrandModel)(nil)

type (
	// ProductBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductBrandModel.
	ProductBrandModel interface {
		productBrandModel
	}

	customProductBrandModel struct {
		*defaultProductBrandModel
	}
)

// NewProductBrandModel returns a model for the database table.
func NewProductBrandModel(conn sqlx.SqlConn, c cache.CacheConf) ProductBrandModel {
	return &customProductBrandModel{
		defaultProductBrandModel: newProductBrandModel(conn, c),
	}
}
