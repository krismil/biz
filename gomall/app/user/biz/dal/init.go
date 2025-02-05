package dal

import (
	"github.com/krismil/biz/gomall/app/user/biz/dal/mysql"
	"github.com/krismil/biz/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
