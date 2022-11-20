package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
用原生go解析并渲染模板
自定义函数传入模板
嵌套模板的使用
2022年11月20日16:06:27
*/

func sayHello(w http.ResponseWriter, r *http.Request) {
	h := func(name []string) (string, error) {
		return name[0] + name[1], nil
	}
	//创建一个模板类的对象
	t := template.New("hello.tmpl")
	//告诉模板现在有一个自定义函数h，在模板中对应的名字是hi
	t.Funcs(template.FuncMap{
		"hi": h,
	})
	//解析模板
	_, err := t.ParseFiles("./test5/hello.tmpl")
	if err != nil {
		fmt.Println("template parse failed,err:", err)
		return
	}
	//渲染模板
	err = t.Execute(w, []string{"新生", "张小搓"})
	if err != nil {
		fmt.Println("template execute failed,err:", err)
		return
	}

}

func demo1(w http.ResponseWriter, r *http.Request) {
	//hello.tmpl为父模板，必须写在第一个
	t, err := template.ParseFiles("./test6/hello.tmpl", "./test6/demo1.tmpl")
	if err != nil {
		fmt.Println("template parse failed,err:", err)
		return
	}

	err = t.Execute(w, "新生张小搓")
	if err != nil {
		fmt.Println("template execute failed,err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/demo", sayHello)
	http.HandleFunc("/", demo1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
