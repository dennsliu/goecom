package logic

import (
	"context"

	"goecom/apps/product/api/internal/svc"
	"goecom/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandsearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandsearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandsearchLogic {
	return &BrandsearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandsearchLogic) Brandsearch(req *types.BrandSearchReq) (resp *types.BrandSearchReply, err error) {
	// todo: add your logic here and delete this line

	return
}
