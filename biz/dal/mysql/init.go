package mysql

import (
	"fmt"

	"github.com/youmie723/GoMall/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	type Version struct {
		Version string
	}
	var v Version
	err := DB.Raw("select version() as version").Scan(&v).Error

	fmt.Println(v)

	if err != nil {
		panic(err)
	}
}
