package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

/*
gin上传文件
2022年11月23日19:02:48
想要上传多个文件时，要在html文件中给file类型的标签加上multiple
处理multipart forms提交文件时默认的内存限制是32 MiB
可以通过下面的方式修改
r.MaxMultipartMemory = 8 << 20  // 8 MiB
*/

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("test13/index.html", "./test13/otherindex.html")

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//上传单个文件
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1") //获取index.html中name=“f1”的文件
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//通过路径与文件名拼接得到文件完整的路径
			dst := path.Join("test13/upload", file.Filename)
			//将上传的文件保存到本地dst路径指向的文件
			err = c.SaveUploadedFile(file, dst)
			msg := "upload successful!"
			if err != nil {
				msg = err.Error()
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": msg,
			})
		}
	})

	r.GET("/uploadMore", func(c *gin.Context) {
		c.HTML(http.StatusOK, "otherindex.html", nil)
	})
	//上传多个文件
	r.POST("/uploadMore", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		} else {
			files := form.File["f2"] //获取otherindex.html中name=“f2”的文件数组
			//循环获取上传的多个文件，重复对单个文件的操作
			for _, file := range files {
				dst := path.Join("test13/upload", file.Filename)
				err = c.SaveUploadedFile(file, dst)
				msg := "upload successful!"
				if err != nil {
					msg = err.Error()
				}
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": msg,
				})
			}
		}
	})

	r.Run(":9090")
}
