package logic

import (
	"context"
	"fmt"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantusergetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantusergetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantusergetLogic {
	return &MerchantusergetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantusergetLogic) Merchantuserget(req *types.MerchantUserGetReq) (resp *types.MerchantUserReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Print("--------------Merchantuserget-------------")
	if req.Id == 0 {
		return nil, errors.Wrapf(ErrUsernamePwdNullError, "Merchantuserget err null id:%d", req.Id)
	}
	userInfo, err := l.svcCtx.MerchantUserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Merchantuserget find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserget find err id:%d, err:%v", req.Id, err)
	}
	return &types.MerchantUserReply{
		Code:        200,
		Msg:         "get merchant user info successfully",
		Id:          userInfo.Id,
		Username:    userInfo.Username,
		Nickname:    userInfo.Nickname,
		Email:       userInfo.Email,
		Telephone:   userInfo.Telephone,
		Mobliephone: userInfo.Mobliephone,
		MerchantId:  userInfo.MerchantId,
		CreatedAt:   userInfo.CreatedAt,
		UpdatedAt:   userInfo.UpdatedAt,
	}, nil
}
