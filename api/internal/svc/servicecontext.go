package svc

import (
	"github.com/blakeyi/gan-circle-server/api/internal/config"
	"github.com/blakeyi/gan-circle-server/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
	}
}
