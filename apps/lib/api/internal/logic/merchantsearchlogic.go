package logic

import (
	"context"
	"fmt"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/pkg/xerr"

	"github.com/pkg/errors"
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
	result, err := l.svcCtx.MerchantModel.Search(req.Keyword, req.Page, req.PageSize, req.OrderBy)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get data merchant list error"), "req: %+v,api err:%+v", req, err)
	}
	size := len(*result)
	fmt.Printf("------------result------size:%d", size)
	var rsp types.MerchantSearchReply
	rsp.Mechants = make([]types.Merchant, size, size)
	for i := 0; i < size; i++ {
		rsp.Mechants[i].Id = (*result)[i].Id
		rsp.Mechants[i].Status = (*result)[i].Status
		rsp.Mechants[i].Name = (*result)[i].Name
	}
	return &rsp, nil
}
