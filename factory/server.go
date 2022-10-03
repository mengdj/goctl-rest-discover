package factory

import (
	"fmt"

	"github.com/mengdj/goctl-rest-discover/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
	"github.com/zeromicro/go-zero/rest"
)

type RestDiscoverServer struct {
	*rest.Server
	config    conf.DiscoverConf
	publisher *discov.Publisher
}

func MustNewServer(c conf.DiscoverConf, opts ...rest.RunOption) *RestDiscoverServer {
	ret := &RestDiscoverServer{
		Server: rest.MustNewServer(c.RestConf, opts...),
		config: c,
	}
	return ret
}

func NewServer(c conf.DiscoverConf, opts ...rest.RunOption) (*RestDiscoverServer, error) {
	ret, err := rest.NewServer(c.RestConf, opts...)
	if nil != err {
		return nil, err
	}
	restDiscoverServer := &RestDiscoverServer{
		Server: ret,
		config: c,
	}
	return restDiscoverServer, nil
}

func (r *RestDiscoverServer) Start() {
	if len(r.config.Etcd.Hosts) > 0 {
		if "0.0.0.0" == r.config.Host {
			r.config.Host = netx.InternalIp()
		}
		r.publisher = discov.NewPublisher(r.config.Etcd.Hosts, r.config.Etcd.Key, fmt.Sprintf("%s:%d", r.config.Host, r.config.Port))
		if err := r.publisher.KeepAlive(); nil != err {
			logx.Errorf("keepalive error:%s", err.Error())
		}
	}
	r.Server.Start()
}

// Stop stops the Server.
func (r *RestDiscoverServer) Stop() {
	if nil != r.publisher {
		r.publisher.Stop()
	}
	r.Server.Stop()
}
