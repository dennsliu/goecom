package svc

import (
	"goecom/apps/lib/api/internal/config"
	"goecom/apps/lib/rpc/lib"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	LibRpc lib.Lib
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		LibRpc: lib.NewLib(zrpc.MustNewClient(c.LibRpc)),
	}
}
