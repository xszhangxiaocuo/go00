/*
Gin为全局路由注册中间件
2022年11月28日19:37:04
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//定义一个中间件，统计程序运行耗时
func mid1() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("mid1 in")
		c.Set("name", "新生张小搓") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		//调用该请求剩余的程序
		c.Next()
		//不调用该请求剩余的程序
		//c.Abort()
		end := time.Since(start) //time.Now().Sub(start)
		fmt.Println("mid1 out")
		fmt.Println(end)
	}

}

func mid2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mid2 in")
		c.Next()
		fmt.Println("mid2 out")
	}
}

func main() {
	//新建一个路由，这个路由没有任何中间件
	r := gin.New()
	//为全部路由注册多个全局中间件
	r.Use(mid1(), mid2())
	r.GET("/", func(c *gin.Context) {
		name := c.MustGet("name").(string) //从上下文取值
		c.JSON(http.StatusOK, gin.H{"msg": name})
	})
	r.Run(":9090")
}
