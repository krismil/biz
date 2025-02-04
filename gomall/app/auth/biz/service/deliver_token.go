package service

import (
	"context"
	auth "github.com/krismil/biz/gomall/rpc_gen/kitex_gen/auth"
)

type DeliverTokenService struct {
	ctx context.Context
} // NewDeliverTokenService new DeliverTokenService
func NewDeliverTokenService(ctx context.Context) *DeliverTokenService {
	return &DeliverTokenService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.

	return
}
