package dal

import (
	"github.com/youmie723/GoMall/biz/dal/mysql"
	//"github.com/youmie723/GoMall/biz/dal/redis"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
