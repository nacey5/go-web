package model

import (
	"fmt"
	_const "go-web/web01_mysql/const"
	"go-web/web01_mysql/util"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试开始,可以在此完成初始化工作")
	util.Init()
	//调用测试链中的其他测试方法
	m.Run()
}

func TestUser_AddUser(t *testing.T) {
	util.Init()
	type fields struct {
		Id       int
		UserName string
		Password string
		Email    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "t1", fields: struct {
			Id       int
			UserName string
			Password string
			Email    string
		}{Id: 0, UserName: "admin", Password: "123456", Email: "admin@qq.com"}, wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &User{
				Id:       tt.fields.Id,
				UserName: tt.fields.UserName,
				Password: tt.fields.Password,
				Email:    tt.fields.Email,
			}
			if err := this.AddUser(); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	defer _const.DB.Close()
}

func TestUser_AddUser2(t *testing.T) {
	util.Init()
	type fields struct {
		Id       int
		UserName string
		Password string
		Email    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "t2", fields: struct {
			Id       int
			UserName string
			Password string
			Email    string
		}{Id: 0, UserName: "admin2", Password: "456789", Email: "admin2@sina.cn"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &User{
				Id:       tt.fields.Id,
				UserName: tt.fields.UserName,
				Password: tt.fields.Password,
				Email:    tt.fields.Email,
			}
			if err := this.AddUser2(); (err != nil) != tt.wantErr {
				t.Errorf("AddUser2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	defer _const.DB.Close()
}

func TestUser_GetUserById(t *testing.T) {
	type fields struct {
		Id       int
		UserName string
		Password string
		Email    string
	}
	//预估结果
	forward := &User{1, "admin", "123456", "admin@qq.com"}
	tests := []struct {
		name     string
		fields   fields
		wantUser *User
		wantErr  bool
	}{
		// TODO: Add test cases.
		{name: "case1", fields: struct {
			Id       int
			UserName string
			Password string
			Email    string
		}{Id: 1, UserName: "admin", Password: "123456", Email: "admin@qq.com"}, wantUser: forward, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &User{
				Id:       tt.fields.Id,
				UserName: tt.fields.UserName,
				Password: tt.fields.Password,
				Email:    tt.fields.Email,
			}
			gotUser, err := this.GetUserById()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(*gotUser)
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUserById() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
	defer _const.DB.Close()
}

func TestUser_GetAllUsers(t *testing.T) {
	type fields struct {
		Id       int
		UserName string
		Password string
		Email    string
	}
	tests := []struct {
		name      string
		fields    fields
		wantUsers []*User
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "case1", fields: struct {
			Id       int
			UserName string
			Password string
			Email    string
		}{}, wantUsers: make([]*User, 2), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &User{
				Id:       tt.fields.Id,
				UserName: tt.fields.UserName,
				Password: tt.fields.Password,
				Email:    tt.fields.Email,
			}
			gotUsers, err := this.GetAllUsers()
			for _, user := range gotUsers {
				if user != nil {
					fmt.Println(*user)
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUsers, tt.wantUsers) {
				t.Errorf("GetAllUsers() gotUsers = %v, want %v", gotUsers, tt.wantUsers)
			}
		})
	}
	defer _const.DB.Close()
}
