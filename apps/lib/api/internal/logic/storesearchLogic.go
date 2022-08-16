package logic

import (
	"context"
	"fmt"
	"math"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/model"
	"goecom/apps/lib/rpc/types/lib"

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
	fmt.Printf("------------result------keyword:%s", req.Keyword)
	var stores []types.Store
	var count int64 = 0
	resultData := l.svcCtx.Gorm.Model(&model.Store{}).Order("id " + req.OrderType)
	var currentPage int64 = 1
	if req.Page > 0 {
		currentPage = req.Page
	}
	resultData.Limit(int(req.PageSize)).Offset(((int(currentPage) - 1) * int(req.PageSize)))
	resultData.Scan(&stores).Limit(-1).Offset(-1).Count(&count)
	err = resultData.Error
	fmt.Printf("------------result------count:%d", count)
	if err != nil {
		fmt.Println(fmt.Sprintf("错误信息:%s", err))
		return nil, err
	}
	fmt.Println(fmt.Sprintf("查询到的条数:%d, %+v", len(stores), stores))
	fmt.Println(fmt.Sprintf("总条数:%d", count))
	size := len(stores)
	fmt.Printf("------------result------size:%d", size)
	var rsp types.StoreSearchReply
	rsp.Code = 200
	rsp.Msg = "Get merchant list successfully"
	rsp.IsEnd = true
	rsp.Stores = make([]types.Store, size)
	for i := 0; i < size; i++ {
		merchantResult, _ := l.svcCtx.LibRpc.GetMerchant(l.ctx, &lib.GetMerchantReq{
			Id: (stores)[i].MerchantId,
		})
		rsp.Stores[i].Id = (stores)[i].Id
		rsp.Stores[i].MerchantId = (stores)[i].MerchantId
		rsp.Stores[i].Merchantname = merchantResult.Name
		rsp.Stores[i].Order = (stores)[i].Order
		rsp.Stores[i].CreatedAt = (stores)[i].CreatedAt
		rsp.Stores[i].UpdatedAt = (stores)[i].UpdatedAt
	}
	rsp.LastVal = (stores)[size-1].Id
	rsp.Total = count
	rsp.CurrentPage = currentPage
	rsp.TotalPage = int64(math.Ceil(float64(count) / float64(req.PageSize)))
	return &rsp, nil
}
