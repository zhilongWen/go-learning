package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// https://topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E8%B7%AF%E7%94%B1/%E4%B8%8A%E4%BC%A0%E5%8D%95%E4%B8%AA%E6%96%87%E4%BB%B6.html

func main() {

	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {

		file, _ := c.FormFile("file")
		//name := c.PostForm("name")

		//c.SaveUploadedFile(file, "./"+file.Filename)
		in, _ := file.Open()
		defer in.Close()

		out, _ := os.Create("./" + file.Filename)
		defer out.Close()

		io.Copy(out, in)

		//c.JSON(200, gin.H{
		//	"status": "success",
		//	"msg":    file,
		//	"name":   name,
		//})

		c.Writer.Header().Add("Content-Disposition", "attachment; filename="+file.Filename)
		c.File("./" + file.Filename)

	})

	r.Run(":8080")
}
