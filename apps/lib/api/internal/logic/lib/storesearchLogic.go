package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoresearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoresearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoresearchLogic {
	return &StoresearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoresearchLogic) Storesearch(req *types.StoreSearchReq) (resp *types.StoreSearchReply, err error) {
	// todo: add your logic here and delete this line

	return
}
