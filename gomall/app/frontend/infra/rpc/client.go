package rpc

import (
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/krismil/biz/gomall/app/frontend/conf"
	frontendUtils "github.com/krismil/biz/gomall/app/frontend/utils"
	"github.com/krismil/biz/gomall/rpc_gen/kitex_gen/user/userservice"
	"sync"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
