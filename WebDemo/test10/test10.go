package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin处理post请求，获取form参数
2022年11月22日21:34:39
*/

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("test10/login.html", "test10/index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		//GetPostForm和DefaultPostForm的用法同GetQuery和DefaultQuery相同
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name":     username,
			"password": password,
		})
	})

	r.Run(":9090")
}
