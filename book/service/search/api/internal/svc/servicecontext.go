package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/book/service/search/api/internal/config"
	"go-zero-demo/book/service/search/api/internal/middleware"
	"go-zero-demo/book/service/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
