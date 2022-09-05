package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoregetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoregetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoregetLogic {
	return &StoregetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoregetLogic) Storeget(req *types.StoreGetReq) (resp *types.StoreReply, err error) {
	// todo: add your logic here and delete this line

	return
}
