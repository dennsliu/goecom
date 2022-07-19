package logic

import (
	"context"

	"goecom/apps/product/api/internal/svc"
	"goecom/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandupdateLogic {
	return &BrandupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandupdateLogic) Brandupdate(req *types.BrandUpdateReq) (resp *types.Brand, err error) {
	// todo: add your logic here and delete this line

	return
}
