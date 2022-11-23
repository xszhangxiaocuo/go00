package main

/*
首次使用gin框架搭建一个服务端
2022年11月15日00:16:12
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认路由
	r := gin.Default()
	//GET：请求方式 "/demo"是请求路径，当浏览器发送GET请求时，会执行后面的匿名函数
	r.GET("/demo", func(c *gin.Context) {
		//返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})

	r.POST("/demo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/demo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/demo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
	//启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":9090")
}
