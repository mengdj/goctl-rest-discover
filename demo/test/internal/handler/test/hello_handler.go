package test

import (
	"net/http"

	"github.com/mengdj/goctl-rest-discover/demo/test/internal/logic/test"
	"github.com/mengdj/goctl-rest-discover/demo/test/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
