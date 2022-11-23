package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin参数绑定
使用ShouldBind
2022年11月23日11:45:19
*/

type UserInfo struct {
	Name string `form:"name" json:"username"` //在tag中用form修改在浏览器输入时的参数名，浏览器发送请求的参数名在tag中决定，默认为结构体中的字段名
	Age  int    `form:"age" json:"age"`       //json的字段名不区分大小写,而form区分大小写
}

func main() {
	r := gin.Default()
	//从querystring中获取数据
	r.GET("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": u.Name,
				"age":  u.Age,
			})
		}
	})
	//从form表单获取数据，也可以从json格式的数据中获取
	r.POST("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": u.Name,
				"age":  u.Age,
			})
		}
	})

	r.Run(":9090")
}
