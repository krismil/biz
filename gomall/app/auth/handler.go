package main

import (
	"context"
	"github.com/krismil/biz/gomall/app/auth/biz/service"
	
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverToken(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewDeliverTokenService(ctx).Run(req)

	return resp, err
}

// VerifyToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyToken(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp, err = service.NewVerifyTokenService(ctx).Run(req)

	return resp, err
}
