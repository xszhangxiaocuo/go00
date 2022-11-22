package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin获取querystring参数
2022年11月22日14:31:28
*/

func main() {
	r := gin.Default()
	//GET请求中URL的 ? 后面的是querystring参数
	r.GET("/", func(c *gin.Context) {
		//通过Query获取浏览器发送的请求中的querystring参数
		name := c.Query("query")                //通过Query获取请求中携带的query参数，如果没有该参数，默认为""
		age := c.DefaultQuery("age", "unknown") //通过Query获取请求中携带的age参数，age参数的默认参数是"unknown"
		//GetQuery会返回一个bool参数，如果浏览器发送的请求中没有msg参数，返回false。这种方法通过ok做判断对msg赋默认值“null”，等效于DefaultQuery方法
		msg, ok := c.GetQuery("msg")
		if !ok {
			msg = "null"
		}
		c.JSON(http.StatusOK, gin.H{
			"name":    name,
			"age":     age,
			"message": msg,
		})
	})
	r.Run(":9090")
}
