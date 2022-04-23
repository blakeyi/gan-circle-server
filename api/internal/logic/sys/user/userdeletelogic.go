package user

import (
	"context"

	"github.com/blakeyi/gan-circle-server/api/internal/svc"
	"github.com/blakeyi/gan-circle-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.UserDeleteReq) (res *types.Response, err error) {
	// todo: add your logic here and delete this line
	// (resp *types.UserInfo, err error)

	return &types.Response{}, nil
}
