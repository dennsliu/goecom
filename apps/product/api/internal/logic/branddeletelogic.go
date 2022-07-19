package logic

import (
	"context"

	"goecom/apps/product/api/internal/svc"
	"goecom/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BranddeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBranddeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BranddeleteLogic {
	return &BranddeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BranddeleteLogic) Branddelete(req *types.BrandDeleteReq) (resp *types.BrandDeleteReply, err error) {
	// todo: add your logic here and delete this line

	return
}
