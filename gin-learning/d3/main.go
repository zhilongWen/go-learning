package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func MiddlewareV1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MiddlewareV1开始执行了")
		c.Next()
		fmt.Println("MiddlewareV1执行完毕")
	}
}

// https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E4%B8%AD%E9%97%B4%E4%BB%B6/%E5%85%A8%E5%B1%80%E4%B8%AD%E9%97%B4%E4%BB%B6.html
func main() {

	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())

	v1 := r.Group("/v1")
	v1.Use(MiddlewareV1())
	v1.GET("/test", func(c *gin.Context) {
		fmt.Println("v1/test")
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
