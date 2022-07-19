package logic

import (
	"context"

	"goecom/apps/product/api/internal/svc"
	"goecom/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandgetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandgetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandgetLogic {
	return &BrandgetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandgetLogic) Brandget(req *types.BrandGetReq) (resp *types.Brand, err error) {
	// todo: add your logic here and delete this line

	return
}
