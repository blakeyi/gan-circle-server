package svc

import (
	"github.com/blakeyi/gan-circle-server/api/internal/config"
	"github.com/blakeyi/gan-circle-server/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"github.com/blakeyi/gan-circle-server/api/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	CheckUrl  rest.Middleware
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	userModel := model.NewUserModel(c.Mongo.Url, c.Mongo.Db,  c.Mongo.Collections.User)
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
		UserModel: userModel,
	}
}
