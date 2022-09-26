package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_const "go-web/web01_mysql/const"
)

func Init() *sql.DB {
	//数据库名:数据库密码@tcp(ip：端口)/数据库名
	_const.DB, _const.Err = sql.Open("mysql", "root:a1160124552@tcp(127.0.0.1:3306)/go_sql")
	//测试文件的时候从外部关闭
	//defer _const.DB.Close()
	if _const.Err != nil {
		panic(_const.Err.Error())
	}
	err := _const.DB.Ping()
	if err != nil {
		panic("链接错误")
	}
	return _const.DB
}
