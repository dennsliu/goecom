package logic

import (
	"context"
	"fmt"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsearchLogic {
	return &MerchantsearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsearchLogic) Merchantsearch(req *types.MerchantSearchReq) (resp *types.MerchantSearchReply, err error) {
	// todo: add your logic here and delete this line

	result, err := l.svcCtx.MerchantModel.FindByKeyword(req.Keyword, req.Page, req.PageSize, req.OrderBy)
	if err != nil {
		return nil, err
	}
	size := len(*result)
	fmt.Printf("------------result------size:%d", size)
	var rsp types.MerchantSearchReply
	rsp.Code = 200
	rsp.Msg = "Get merchant list successfully"
	rsp.IsEnd = true

	rsp.Mechants = make([]types.Merchant, size, size)
	for i := 0; i < size; i++ {
		rsp.Mechants[i].Id = (*result)[i].Id
		rsp.Mechants[i].Name = (*result)[i].Name
		rsp.Mechants[i].Name = (*result)[i].Name
		rsp.Mechants[i].Status = (*result)[i].Status
		rsp.Mechants[i].CreatedAt = (*result)[i].CreatedAt
		rsp.Mechants[i].UpdatedAt = (*result)[i].UpdatedAt
	}
	rsp.LastVal = (*result)[size-1].Id
	return &rsp, nil
}
