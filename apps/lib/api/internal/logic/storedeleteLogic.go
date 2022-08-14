package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoredeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoredeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoredeleteLogic {
	return &StoredeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoredeleteLogic) Storedelete(req *types.StoreDeleteReq) (resp *types.StoreDeleteReply, err error) {
	// todo: add your logic here and delete this line

	return
}
