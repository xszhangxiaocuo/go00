package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin实现http请求重定向
2022年11月23日20:48:23
*/

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://81.68.219.45")
	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //修改请求的URI地址，将请求转向路由b
		r.HandleContext(c)        //继续执行后面的操作
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "this is /b",
		})
	})
	r.Run(":9090")
}
