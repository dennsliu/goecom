package logic

import (
	"context"

	"goecom/apps/product/api/internal/svc"
	"goecom/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandaddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandaddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandaddLogic {
	return &BrandaddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandaddLogic) Brandadd(req *types.BrandAddReq) (resp *types.Brand, err error) {
	// todo: add your logic here and delete this line

	return
}
