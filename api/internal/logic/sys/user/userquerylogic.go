package user

import (
	"context"

	"github.com/blakeyi/gan-circle-server/api/internal/svc"
	"github.com/blakeyi/gan-circle-server/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryLogic {
	return &UserQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserQueryLogic) UserQuery(req *types.UserQueryReq) (res *types.Response, err error) {
	// todo: add your logic here and delete this line
	// (resp *types.UserInfo, err error)
	logx.Info("UserQuery")
	res = &types.Response{}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, "619e599722f04d18f1b1e8af")
	if err != nil {
		logx.Error(err.Error())
		return res, err
	}
	res.SetErrorCode(types.Succeed)
	res.RetContent = user
	return res, nil
}
