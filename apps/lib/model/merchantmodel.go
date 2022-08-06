package model

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MerchantModel = (*customMerchantModel)(nil)

type (
	// MerchantModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMerchantModel.
	MerchantModel interface {
		merchantModel
		FindByKeyword(keyword string, page int64, pageSize int64, orderBy string) (*[]Merchant, error)
	}

	customMerchantModel struct {
		*defaultMerchantModel
	}
)

// NewMerchantModel returns a model for the database table.
func NewMerchantModel(conn sqlx.SqlConn, c cache.CacheConf) MerchantModel {
	return &customMerchantModel{
		defaultMerchantModel: newMerchantModel(conn, c),
	}
}

func (m *defaultMerchantModel) FindByKeyword(keyword string, page int64, pageSize int64, orderBy string) (*[]Merchant, error) {

	if orderBy == "" {
		orderBy = "id"
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	var resp []Merchant
	query := fmt.Sprintf("select %s from %s order by %s limit %d", merchantRows, m.table, orderBy, pageSize)
	err := m.QueryRowsNoCache(&resp, query)
	fmt.Println(err)
	fmt.Println(resp)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}
