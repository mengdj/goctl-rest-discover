package svc

import (
	"github.com/mengdj/goctl-rest-discover/demo/exa2/client"
	"github.com/mengdj/goctl-rest-discover/demo/test/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	HelloClient client.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		HelloClient: client.MustClient(c.HelloDiscoverConf),
	}
}
