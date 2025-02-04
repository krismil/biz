package service

import (
	"context"
	auth "github.com/krismil/biz/gomall/rpc_gen/kitex_gen/auth"
	"testing"
)

func TestVerifyToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyTokenService(ctx)
	// init req and assert value

	req := &auth.VerifyTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
