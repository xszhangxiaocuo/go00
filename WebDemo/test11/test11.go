package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin获取请求的path(URI)参数
2022年11月22日23:34:55
*/

func main() {
	r := gin.Default()
	//冒号后的即为参数名
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")

		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	r.Run(":9090")
}
