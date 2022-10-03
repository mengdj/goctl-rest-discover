package test

import (
	"context"

	"github.com/mengdj/goctl-rest-discover/demo/exa2/client"
	"github.com/mengdj/goctl-rest-discover/demo/test/internal/svc"
	"github.com/mengdj/goctl-rest-discover/demo/test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello() (*types.Response, error) {
	var (
		ret = &types.Response{
			Code: 0,
			Msg:  "",
		}
		err  error = nil
		resp *client.Response
	)
	resp, err = l.svcCtx.HelloClient.Hello(l.ctx, &client.HelloRequest{
		Msg: "hello,rest",
	})
	if nil != err {
		ret.Code = 1
		ret.Msg = err.Error()
	} else {
		ret.Code = resp.Code
		ret.Msg = resp.Msg
	}
	return ret, nil
}