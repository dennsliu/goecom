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

type LanguagessearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLanguagessearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LanguagessearchLogic {
	return &LanguagessearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LanguagessearchLogic) Languagessearch(req *types.LanguagesSearchReq) (resp *types.LanguagesSearchReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("------------result------keyword:%s", req.Keyword)
	var languages []types.Languages
	var count int64 = 0
	resultData := l.svcCtx.Gorm.Model(&model.Languages{}).Order("id " + req.OrderType)
	if len(req.Keyword) > 0 {
		resultData.Where("name like ?", "%"+req.Keyword+"%")
	}
	var currentPage int64 = 1
	if req.Page > 0 {
		currentPage = req.Page
	}
	resultData.Limit(int(req.PageSize)).Offset(((int(currentPage) - 1) * int(req.PageSize)))

	resultData.Scan(&languages).Limit(-1).Offset(-1).Count(&count)
	err = resultData.Error
	fmt.Printf("------------result------count:%d", count)
	if err != nil {
		fmt.Println(fmt.Sprintf("错误信息:%s", err))
		return nil, err
	}
	fmt.Println(fmt.Sprintf("查询到的条数:%d, %+v", len(languages), languages))
	fmt.Println(fmt.Sprintf("总条数:%d", count))
	size := len(languages)
	fmt.Printf("------------result------size:%d", size)
	var rsp types.LanguagesSearchReply
	rsp.Code = 200
	rsp.Msg = "Get merchant list successfully"
	rsp.IsEnd = true

	rsp.Languages = make([]types.Languages, size)
	for i := 0; i < size; i++ {
		rsp.Languages[i].Id = (languages)[i].Id
		rsp.Languages[i].Name = (languages)[i].Name
		rsp.Languages[i].Code = (languages)[i].Code
		rsp.Languages[i].Image = (languages)[i].Image
		rsp.Languages[i].Directory = (languages)[i].Directory
		rsp.Languages[i].Order = (languages)[i].Order
		rsp.Languages[i].CreatedAt = (languages)[i].CreatedAt
		rsp.Languages[i].UpdatedAt = (languages)[i].UpdatedAt
	}
	rsp.LastVal = (languages)[size-1].Id
	rsp.Total = count
	rsp.CurrentPage = currentPage
	rsp.TotalPage = int64(math.Ceil(float64(count) / float64(req.PageSize)))
	return &rsp, nil
}
