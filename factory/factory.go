package factory

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"

	"github.com/mengdj/goctl-rest-discover/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/rest/httpc"
)

type (
	RestDiscoverFactoryOption func(*RestDiscoverFactory)
	RestDiscoverFactory       struct {
		config   conf.DiscoverClientConf
		protocol string
		base     []string
		service  httpc.Service
		rwMutex  sync.RWMutex
	}
)

func NewRestDiscoverFactory(c conf.DiscoverClientConf, opts ...RestDiscoverFactoryOption) *RestDiscoverFactory {
	ret := &RestDiscoverFactory{
		protocol: "http://",
		config:   c,
		service:  httpc.NewService(c.Etcd.Key),
	}
	for _, opt := range opts {
		opt(ret)
	}
	//https
	if c.TLS {
		//use https
		ret.protocol = "https://"
	}
	if len(c.Endpoints) > 0 {
		ret.base = c.Endpoints
	} else {
		opts := make([]discov.SubOption, 0)
		if "" != c.Etcd.User {
			opts = append(opts, discov.WithSubEtcdAccount(c.Etcd.User, c.Etcd.Pass))
		}
		sub, err := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key, opts...)
		if nil != err {
			panic(err)
		}
		//get base address
		update := func() {
			ret.rwMutex.Lock()
			defer func() {
				ret.rwMutex.Unlock()
			}()
			if values := sub.Values(); len(values) > 0 {
				ret.base = values
			}
		}
		sub.AddListener(update)
		update()
	}
	return ret
}

func (f *RestDiscoverFactory) getBase() string {
	f.rwMutex.RLock()
	defer func() {
		f.rwMutex.RUnlock()
	}()
	if len(f.base) > 0 {
		rand.Shuffle(len(f.base), func(i, j int) {
			f.base[i], f.base[j] = f.base[j], f.base[i]
		})
		return f.base[0]
	}
	return ""
}

func (f *RestDiscoverFactory) Invoke(ctx context.Context, method string, path string, data interface{}, resp interface{}) error {
	var (
		base   = f.getBase()
		result *http.Response
		err    error
	)
	if "" == base {
		return errors.New("base can't nil")
	}
	result, err = f.service.Do(ctx, strings.ToUpper(method), fmt.Sprintf("%s%s%s", f.protocol, f.getBase(), path), data)
	if nil != err {
		return err
	}
	if http.StatusOK != result.StatusCode {
		//must 200
		return errors.New(result.Status)
	}
	if nil != data {
		if err = jsonx.UnmarshalFromReader(result.Body, resp); nil != err {
			return err
		}
	}
	return nil
}
