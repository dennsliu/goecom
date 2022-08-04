package svc

import (
	"goecom/apps/lib/rpc/internal/config"

	"goecom/apps/lib/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	RedisClient       *redis.Redis
	MerchantModel     model.MerchantModel
	MerchantUserModel model.MerchantUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			//		r.Pass = c.Redis.Pass
		}),
		MerchantModel:     model.NewMerchantModel(sqlConn, c.Cache),
		MerchantUserModel: model.NewMerchantUserModel(sqlConn, c.Cache),
	}
}
