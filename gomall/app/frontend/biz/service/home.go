package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	home "github.com/krismil/biz/gomall/app/frontend/hertz_gen/hertz/home"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "1号", "Price": 1.0, "Picture": "/static/image/logo.png"},
		{"Name": "2号", "Price": 1.0, "Picture": "/static/image/logo.png"},
		{"Name": "1号", "Price": 1.0, "Picture": "/static/image/home.png"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
