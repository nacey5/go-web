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

// GetUserById 获取单条的User信息
func (this *User) GetUserById() (user *User, err error) {
	sqlStr := "select id,user_name,password,email from users where id=?"
	//执行
	row := _const.DB.QueryRow(sqlStr, this.Id)
	var id int
	var userName string
	var password string
	var email string
	err = row.Scan(&id, &userName, &password, &email)
	if err != nil {
		return nil, err
	}
	user = &User{
		id, userName, password, email,
	}
	return
}

// GetAllUsers 获取所有的User
func (this *User) GetAllUsers() (users []*User, err error) {
	sqlStr := "select id,user_name,password,email from users where 1=1"
	rows, err := _const.DB.Query(sqlStr)
	if err != nil {
		return
	}
	users = make([]*User, 2)
	for rows.Next() {
		var id int
		var userName string
		var password string
		var email string
		err = rows.Scan(&id, &userName, &password, &email)
		if err != nil {
			return
		}
		user := &User{
			id, userName, password, email,
		}
		users = append(users, user)
	}
	return
}
