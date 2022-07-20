package svc

import (
	"goecom/apps/lib/api/internal/config"
	"goecom/apps/lib/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	MerchantModel     model.MerchantModel
	MerchantUserModel model.MerchantUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		MerchantModel:     model.NewMerchantModel(conn, c.CacheRedis),
		MerchantUserModel: model.NewMerchantUserModel(conn, c.CacheRedis),
	}
}
