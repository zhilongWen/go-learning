package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hello, world!")

	r := gin.Default()

	// http://127.0.0.1:8001/path/123&username=admin&password=123456    => id  = 123
	r.GET("/path/:id", func(c *gin.Context) {

		id := c.Param("id")
		username := c.DefaultQuery("username", "admin")
		password := c.Query("password")

		c.JSON(200, gin.H{
			"success":  true,
			"id":       id,
			"username": username,
			"password": password,
		})
	})

	r.POST("/path/", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(200, gin.H{
			"success":  true,
			"username": username,
			"password": password,
		})
	})

	r.PUT("/path/", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(200, gin.H{
			"success":  true,
			"username": username,
			"password": password,
		})
	})

	r.DELETE("/path/:id", func(c *gin.Context) {

		id := c.Param("id")
		c.JSON(200, gin.H{
			"success": true,
			"id":      id,
		})
	})

	err := r.Run(":8001")
	if err != nil {
		fmt.Println("Error starting admin-be ", err)
		return
	}
}
