package conf

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type DiscoverConf struct {
	rest.RestConf
	Etcd      discov.EtcdConf `json:",optional"` //服务中心
	Endpoints []string        `json:",optional"` //直连
	TLS       bool            `json:"tls,optional"`
}
