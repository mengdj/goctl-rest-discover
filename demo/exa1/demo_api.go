package main

import (
	"flag"
	"fmt"

	"github.com/mengdj/goctl-rest-discover/demo/exa1/internal/config"
	"github.com/mengdj/goctl-rest-discover/demo/exa1/internal/handler"
	"github.com/mengdj/goctl-rest-discover/demo/exa1/internal/svc"
	"github.com/mengdj/goctl-rest-discover/factory"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/demo_api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := factory.MustNewServer(c.DiscoverConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server.Server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
