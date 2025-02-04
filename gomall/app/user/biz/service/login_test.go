package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/krismil/biz/gomall/app/user/biz/dal/mysql"
	user "github.com/krismil/biz/gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "zhangsan@163.com",
		Password: "12345",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
