package logic

import (
	"context"

	"go-zero-demo/book/service/user/rpc/internal/svc"
	"go-zero-demo/book/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	// 添加逻辑
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:     userInfo.Id,
		Name:   userInfo.Name,
		Number: userInfo.Name,
		Gender: userInfo.Gender,
	}, nil
}
