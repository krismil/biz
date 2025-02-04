package dal

import (
	"github.com/krismil/biz/gomall/app/user/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
