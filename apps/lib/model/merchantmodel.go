package model

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MerchantModel = (*customMerchantModel)(nil)

type (
	// MerchantModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMerchantModel.
	MerchantModel interface {
		merchantModel
		FindByKeyword(keyword string, status int64, page int64, pageSize int64, lastId int64, orderType string) (*[]Merchant, error)
		FindByKeywordCount(ctx context.Context, keyword string, status int64) (int64, error)
		FindOneOrderById(orderType string) (*Merchant, error)
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

func (m *defaultMerchantModel) FindByKeyword(keyword string, status int64, page int64, pageSize int64, lastId int64, orderType string) (*[]Merchant, error) {
	fmt.Printf("------------FindByKeyword------keyword:%s", keyword)
	if len(orderType) == 0 {
		orderType = "desc"
	}
	var orderTypeString string
	if orderType == "asc" {
		orderTypeString = "id>" + strconv.FormatInt(lastId, 10)
	} else {
		orderTypeString = "id<" + strconv.FormatInt(lastId, 10)
	}
	fmt.Printf("------------FindByKeyword------orderTypeString:%s", orderTypeString)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	var resp []Merchant
	if len(keyword) > 0 {
		query := fmt.Sprintf("select %s from %s where name like %s and status=%d and %s order by id %s limit %d", merchantRows, m.table, "'%"+keyword+"%'", status, orderTypeString, orderType, pageSize)
		err := m.QueryRowsNoCache(&resp, query)
		fmt.Println(err)
		fmt.Println(resp)
		switch err {
		case nil:
			return &resp, nil
		default:
			return nil, err
		}
	} else {
		query := fmt.Sprintf("select %s from %s where %s and status=%d order by id %s limit %d", merchantRows, m.table, orderTypeString, status, orderType, pageSize)
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
}
func (m *defaultMerchantModel) FindByKeywordCount(ctx context.Context, keyword string, status int64) (int64, error) {
	resp := status
	fmt.Printf("------------FindByKeywordCount------keyword:%s", keyword)
	if len(keyword) > 0 {
		fmt.Printf("select %s from %s where name like %s and status=%d", "count(id) as total", m.table, "'%"+keyword+"%'", status)
		query := fmt.Sprintf("select %s from %s where name like %s and status=%d", "count(id) as total", m.table, "'%"+keyword+"%'", status)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
		fmt.Println(err)
		fmt.Println(&resp)
		fmt.Printf("------------FindByKeywordCount------resp:%d", resp)
		switch err {
		case nil:
			return resp, nil
		default:
			return 0, err
		}
	} else {
		fmt.Printf("select %s from %s where status=%d", "count(id) as total", m.table, status)
		query := fmt.Sprintf("select %s from %s where status=%d", "count(id) as total", m.table, status)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
		fmt.Println(err)
		fmt.Println(&resp)
		fmt.Printf("------------FindByKeywordCount------resp:%d", resp)
		switch err {
		case nil:
			return resp, nil
		default:
			return 0, err
		}
	}
}
func (m *defaultMerchantModel) FindOneOrderById(orderType string) (*Merchant, error) {
	var resp Merchant
	query := fmt.Sprintf("select %s from %s order by id %s limit 1", merchantRows, m.table, orderType)
	err := m.QueryRowsNoCache(&resp, query)
	fmt.Println(err)
	fmt.Println(resp)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
