/*
Gin中间件
2022年11月27日21:22:02
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
		//调用该请求剩余的程序
		c.Next()
		//不调用该请求剩余的程序
		//c.Abort()
		end := time.Since(start) //time.Now().Sub(start)
		fmt.Println("mid1 out")
		fmt.Println(end)
	}

}

func main() {
	r := gin.Default()
	r.GET("/", mid1(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
	r.Run(":9090")
}
