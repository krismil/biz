package auth

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	auth "github.com/krismil/biz/gomall/rpc_gen/kitex_gen/auth"
)

func DeliverToken(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (resp *auth.DeliveryResp, err error) {
	resp, err = defaultClient.DeliverToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeliverToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyToken(ctx context.Context, req *auth.VerifyTokenReq, callOptions ...callopt.Option) (resp *auth.VerifyResp, err error) {
	resp, err = defaultClient.VerifyToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
