package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板，必须要有html
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "Hello T")
}
