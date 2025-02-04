package auth

import (
	"context"
	auth "github.com/krismil/biz/gomall/rpc_gen/kitex_gen/auth"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	
)

type RPCClient interface {
	KitexClient() authservice.Client
	Service() string
	DeliverToken(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
	VerifyToken(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := authservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient authservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() authservice.Client {
	return c.kitexClient
}

func (c *clientImpl) DeliverToken(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	return c.kitexClient.DeliverToken(ctx, Req, callOptions...)
}

func (c *clientImpl) VerifyToken(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error) {
	return c.kitexClient.VerifyToken(ctx, Req, callOptions...)
}
