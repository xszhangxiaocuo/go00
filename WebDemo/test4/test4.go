package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
用原生go解析并渲染模板
向模板中同时传入结构体，map类型的数据
2022年11月15日22:27:43
*/

type User struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./test4/hello.tmpl")
	if err != nil {
		fmt.Println("template parse failed,err:", err)
		return
	}
	user := User{
		Name:   "新生张小搓",
		Age:    19,
		Gender: "man",
	}
	m := map[string]interface{}{
		"name":   "新生张小搓",
		"age":    19,
		"gender": "man",
	}
	//渲染模板
	err = t.Execute(w, map[string]interface{}{
		"user": user,
		"m":    m,
	})
	if err != nil {
		fmt.Println("template execute failed,err:", err)
		return
	}

}

func main() {
	http.HandleFunc("/demo", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
