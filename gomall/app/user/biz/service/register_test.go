package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/krismil/biz/gomall/app/user/biz/dal/mysql"
	user "github.com/krismil/biz/gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "zhangsan@163.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
