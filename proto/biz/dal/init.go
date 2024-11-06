package dal

import (
	"github.com/youmei723/GoMall/biz/dal/mysql"
	"github.com/youmei723/GoMall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
