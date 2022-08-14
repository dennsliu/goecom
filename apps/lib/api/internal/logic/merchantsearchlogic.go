package logic

import (
	"context"
	"fmt"
	"math"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/model"

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
	fmt.Printf("------------result------keyword:%s", req.Keyword)
	var merchants []types.Merchant
	var count int64 = 0
	resultData := l.svcCtx.Gorm.Model(&model.Merchant{}).Where("status=?", req.Status).Order("id " + req.OrderType)
	if len(req.Keyword) > 0 {
		resultData.Where("name like ?", "%"+req.Keyword+"%")
	}
	var currentPage int64 = 1
	if req.Page > 0 {
		currentPage = req.Page
	}
	resultData.Limit(int(req.PageSize)).Offset(((int(currentPage) - 1) * int(req.PageSize)))

	resultData.Scan(&merchants).Limit(-1).Offset(-1).Count(&count)
	err = resultData.Error
	fmt.Printf("------------result------count:%d", count)
	if err != nil {
		fmt.Println(fmt.Sprintf("错误信息:%s", err))
		return nil, err
	}
	fmt.Println(fmt.Sprintf("查询到的条数:%d, %+v", len(merchants), merchants))
	fmt.Println(fmt.Sprintf("总条数:%d", count))
	size := len(merchants)
	fmt.Printf("------------result------size:%d", size)
	var rsp types.MerchantSearchReply
	rsp.Code = 200
	rsp.Msg = "Get merchant list successfully"
	rsp.IsEnd = true

	rsp.Mechants = make([]types.Merchant, size)
	for i := 0; i < size; i++ {
		rsp.Mechants[i].Id = (merchants)[i].Id
		rsp.Mechants[i].Name = (merchants)[i].Name
		rsp.Mechants[i].Status = (merchants)[i].Status
		rsp.Mechants[i].CreatedAt = (merchants)[i].CreatedAt
		rsp.Mechants[i].UpdatedAt = (merchants)[i].UpdatedAt
	}
	rsp.LastVal = (merchants)[size-1].Id
	rsp.Total = count
	rsp.CurrentPage = currentPage
	rsp.TotalPage = int64(math.Ceil(float64(count) / float64(req.PageSize)))
	return &rsp, nil
}
