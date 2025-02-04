package dal

import (
	"github.com/krismil/biz/gomall/app/auth/biz/dal/mysql"
	"github.com/krismil/biz/gomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
