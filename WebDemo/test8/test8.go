package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
用gin框架返回json数据
2022年11月21日23:28:50
*/

func main() {
	r := gin.Default()
	//方法一：利用gin.H直接传入json数据
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "新生张小搓",
			"age":     19,
			"message": "hello",
		})
	})
	//方法二：自定义结构体传入数据
	type msg struct {
		Name    string `json:"name"` //用tag对结构体字段做定制化操作，修改其在json中的变量名
		Age     int
		Message string
	}

	data := msg{
		Name:    "新生张小搓",
		Age:     19,
		Message: "Hello",
	}
	r.GET("/anotherJson", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9090")
}
