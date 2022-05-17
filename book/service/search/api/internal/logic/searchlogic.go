package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-demo/book/service/user/rpc/user"

	"go-zero-demo/book/service/search/api/internal/svc"
	"go-zero-demo/book/service/search/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (*types.SearchReply, error) {
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	// 使用user rpc
	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SearchReply{
		Name:  userInfo.Name,
		Count: 100,
	}, nil
}
