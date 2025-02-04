package service

import (
	"context"
	auth "github.com/krismil/biz/gomall/rpc_gen/kitex_gen/auth"
	"testing"
)

func TestDeliverToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeliverTokenService(ctx)
	// init req and assert value

	req := &auth.DeliverTokenReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
