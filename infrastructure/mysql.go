package infrastructure

import (
	"fmt"
	"xorm.io/xorm"
)

var MysqlDB *xorm.Engine = initMysqlDB()

var (
	MYSQL_HOST    = "tcp(13.231.174.2:3306)"
	MYSQL_DB      = "talentpool"
	MYSQL_SECRECT = "yYVim5WbqzkWziNY"
)

func initMysqlDB() *xorm.Engine {
	//var MYSQL_HOST = "tcp(127.0.0.1:3306)"
	//var MYSQL_SECRECT = "123"
	//EEpLWKlYixYtYGSx
	//var MYSQL_HOST = "tcp(192.168.1.165:3306)"
	//MYSQL_SECRECT = "a"
	//var MYSQL_DB = "talentpool"
	//var DEEPL_FREE_API_KEY = "ed8fb40e-858f-7167-44c8-65ec333131c2:fx"

	db, _ := xorm.NewEngine("mysql", fmt.Sprintf("root:%s@%s/%s?charset=utf8", MYSQL_SECRECT, MYSQL_HOST, MYSQL_DB))
	db.ShowSQL(true)
	return db
}
