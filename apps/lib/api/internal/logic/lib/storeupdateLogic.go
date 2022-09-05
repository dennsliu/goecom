package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreupdateLogic {
	return &StoreupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreupdateLogic) Storeupdate(req *types.StoreUpdateReq) (resp *types.StoreReply, err error) {
	// todo: add your logic here and delete this line

	return
}
