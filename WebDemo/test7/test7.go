package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
用gin返回一个网页模板
2022年11月21日00:36:09
*/

func main() {
	r := gin.Default()
	r.Static("/xxx", "./test7/static")
	r.LoadHTMLGlob("test7/templates/**/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	r.GET("/faq", func(c *gin.Context) {
		c.HTML(http.StatusOK, "faq.html", nil)
	})

	r.GET("/service", func(c *gin.Context) {
		c.HTML(http.StatusOK, "service.html", nil)
	})

	r.GET("/single-service", func(c *gin.Context) {
		c.HTML(http.StatusOK, "single-service.html", nil)
	})

	r.GET("/team", func(c *gin.Context) {
		c.HTML(http.StatusOK, "team.html", nil)
	})
	r.Run(":9090")
}
