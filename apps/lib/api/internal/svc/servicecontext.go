package svc

import (
	"goecom/apps/lib/api/internal/config"
	"goecom/apps/lib/model"
	"goecom/apps/lib/rpc/lib"
	"goecom/pkg/xgorm"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Redis
	LibRpc            lib.Lib
	MerchantModel     model.MerchantModel
	MerchantUserModel model.MerchantUserModel
	Gorm              *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:            c,
		LibRpc:            lib.NewLib(zrpc.MustNewClient(c.LibRpc)),
		MerchantModel:     model.NewMerchantModel(sqlConn, c.Cache),
		MerchantUserModel: model.NewMerchantUserModel(sqlConn, c.Cache),
		Gorm:              xgorm.NewGorm(c.DB.DataSource),
	}
}
