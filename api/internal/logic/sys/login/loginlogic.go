package login

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/blakeyi/gan-circle-server/api/internal/svc"
	"github.com/blakeyi/gan-circle-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (res *types.Response, err error) {
	// todo: add your logic here and delete this line
	// (resp *types.LoginRes, err error)
	res = &types.Response{}
	data := &types.LoginRes{}

	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, req.Openid+req.Phone,
		l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Error(err.Error())
		return res, err
	}
	data.AccessToken = token
	data.AccessExpire = time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire
	data.RefreshAfter = time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire/2
	data.Uid = 1111
	res.RetContent = data
	res.SetErrorCode(types.Succeed)
	return res, nil
}

func (l *LoginLogic) getJwtToken(secretKey, userInfo string, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = seconds
	claims["userInfo"] = userInfo
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
