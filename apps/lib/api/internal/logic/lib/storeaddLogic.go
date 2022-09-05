package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreaddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreaddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreaddLogic {
	return &StoreaddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreaddLogic) Storeadd(req *types.StoreAddReq) (resp *types.StoreReply, err error) {
	// todo: add your logic here and delete this line

	return
}
