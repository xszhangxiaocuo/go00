package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
用原生go解析并渲染模板
2022年11月15日22:27:43
*/

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./test3/hello.tmpl")
	if err != nil {
		fmt.Println("template parse failed,err:", err)
		return
	}
	//渲染模板
	err = t.Execute(w, "新生张小搓")
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
