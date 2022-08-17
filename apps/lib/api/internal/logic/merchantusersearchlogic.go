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

type MerchantusersearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantusersearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantusersearchLogic {
	return &MerchantusersearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantusersearchLogic) Merchantusersearch(req *types.MerchantSearchReq) (resp *types.MerchantUserSearchReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("------------result------keyword:%s", req.Keyword)
	var merchantUsers []types.MerchantUser
	var count int64 = 0
	resultData := l.svcCtx.Gorm.Model(&model.MerchantUser{}).Where("status=?", req.Status).Order("id " + req.OrderType)
	if len(req.Keyword) > 0 {
		resultData.Where("username like ?", "%"+req.Keyword+"%")
	}
	var currentPage int64 = 1
	if req.Page > 0 {
		currentPage = req.Page
	}
	resultData.Limit(int(req.PageSize)).Offset(((int(currentPage) - 1) * int(req.PageSize)))

	resultData.Scan(&merchantUsers).Limit(-1).Offset(-1).Count(&count)
	err = resultData.Error
	fmt.Printf("------------result------count:%d", count)
	if err != nil {
		fmt.Println(fmt.Sprintf("错误信息:%s", err))
		return nil, err
	}
	fmt.Println(fmt.Sprintf("查询到的条数:%d, %+v", len(merchantUsers), merchantUsers))
	fmt.Println(fmt.Sprintf("总条数:%d", count))
	size := len(merchantUsers)
	fmt.Printf("------------result------size:%d", size)
	var rsp types.MerchantUserSearchReply
	rsp.Code = 200
	rsp.Msg = "Get merchant list successfully"
	rsp.IsEnd = true

	rsp.MechantUsers = make([]types.MerchantUser, size)
	for i := 0; i < size; i++ {
		merchantResult, err := l.svcCtx.LibRpc.GetMerchant(l.ctx, &lib.GetMerchantReq{
			Id: (merchantUsers)[i].MerchantId,
		})
		var merchantname string
		if err != nil {
			merchantname = ""
		} else {
			merchantname = merchantResult.Name
		}
		rsp.MechantUsers[i].Id = (merchantUsers)[i].Id
		rsp.MechantUsers[i].Nickname = (merchantUsers)[i].Nickname
		rsp.MechantUsers[i].Username = (merchantUsers)[i].Username
		rsp.MechantUsers[i].Email = (merchantUsers)[i].Email
		rsp.MechantUsers[i].Telephone = (merchantUsers)[i].Telephone
		rsp.MechantUsers[i].Mobliephone = (merchantUsers)[i].Mobliephone
		rsp.MechantUsers[i].MerchantId = (merchantUsers)[i].MerchantId
		rsp.MechantUsers[i].Merchantname = merchantname
		rsp.MechantUsers[i].Status = (merchantUsers)[i].Status
		rsp.MechantUsers[i].CreatedAt = (merchantUsers)[i].CreatedAt
		rsp.MechantUsers[i].UpdatedAt = (merchantUsers)[i].UpdatedAt
	}
	rsp.LastVal = (merchantUsers)[size-1].Id
	rsp.Total = count
	rsp.CurrentPage = currentPage
	rsp.TotalPage = int64(math.Ceil(float64(count) / float64(req.PageSize)))
	return &rsp, nil
}
