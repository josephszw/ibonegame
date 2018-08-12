package mysql

import (
	"ibgame/logs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func GetEngine() (ret *xorm.Engine, err error) {

	if ret, err = xorm.NewEngine("mysql", "work:123456@tcp(39.107.94.42:3306)/go?charset=utf8"); err == nil {
	} else {
		logs.Error.Println("db :err ", err)
	}
	return
}
