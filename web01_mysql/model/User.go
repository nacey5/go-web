package model

import (
	"fmt"
	_const "go-web/web01_mysql/const"
)

// User User结构体
type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (this *User) AddUser() (err error) {
	sqlStr := "insert into users(user_name,password,email) values (?,?,?)"
	//预编译
	inStmt, err := _const.DB.Prepare(sqlStr)
	if err != nil {
		//panic("编译错误")
		fmt.Println("编译出现异常")
		return
	}
	//执行
	_, err = inStmt.Exec(this.UserName, this.Password, this.Email)
	if err != nil {
		panic("执行出错")
	}
	return nil
}

// AddUser2 不带Prepare的方法
func (this *User) AddUser2() (err error) {
	sqlStr := "insert into users(user_name,password,email) values (?,?,?)"
	_, err = _const.DB.Exec(sqlStr, this.UserName, this.Password, this.Email)
	if err != nil {
		panic("编译错误")
		return
	}
	return
}
