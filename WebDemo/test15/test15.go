package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin路由和路由组
路由组多用于区分不同的业务线或API版本
2022年11月23日21:18:01
*/

func main() {
	r := gin.Default()
	//当访问的路由不存在时，自定义404页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "该网页不存在！",
		})
	})
	//将公用的前缀提取出来作为路由组，即访问的路由为"/index/xxx"
	group := r.Group("/index")
	{ //用代码块将路由组的代码放在一起，较为整洁
		group.GET("/index1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "index1"})
		})

		group.GET("/index2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "index2"})
		})

		group.GET("/index3", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "index3"})
		})
	}
	r.Run(":9090")
}
