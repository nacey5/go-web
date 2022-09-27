package main

import (
	"encoding/json"
	"fmt"
	"go-web/web01_mysql/model"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testJson2", testJsonRes2)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "您发送的请求地址是:", r.URL.Path)
	fmt.Fprintln(w, "您发送的请求地址后的查询字符是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头的信息有:", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头的信息有:", r.Header.Get("Accept-Encoding"))
	//获取请求体中的内容的长度
	//length := r.ContentLength
	//body := make([]byte, length)
	//将请求体中的内容读到body中
	//r.Body.Read(body)
	//在浏览器中显示请求体中的内容
	//fmt.Fprintln(w, "请求体中的内容有:\n", string(body))

	//解析表单,在调用r.Form之前必须进行这操作
	r.ParseForm()
	//获取请求参数
	fmt.Fprintln(w, "请求参数有:\n", r.PostForm)
	fmt.Fprintln(w, "URL中的User请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "URL中的UserName请求参数的值是：", r.PostFormValue("name"))
}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	//设置响应的内容的类型
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "https://www.baidu.com")
	user := model.User{
		Id:       1,
		UserName: "zs",
		Password: "123456",
		Email:    "admin@qq.com",
	}
	data, _ := json.Marshal(user)
	w.Write(data)
}

func testJsonRes2(w http.ResponseWriter, r *http.Request) {
	//设置响应的内容的类型
	w.Header().Set("Location", "https://www.baidu.com")
	//设置响应状态码
	w.WriteHeader(302)
}
