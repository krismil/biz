package service

import (
	"context"
	auth "github.com/krismil/biz/gomall/rpc_gen/kitex_gen/auth"
)

type VerifyTokenService struct {
	ctx context.Context
} // NewVerifyTokenService new VerifyTokenService
func NewVerifyTokenService(ctx context.Context) *VerifyTokenService {
	return &VerifyTokenService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.

	return
}
