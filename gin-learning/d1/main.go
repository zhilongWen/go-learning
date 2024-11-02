package main

import "github.com/gin-gonic/gin"

type User struct {
	Name     string `json:"name" uri:"name" yaml:"name" form:"name" binding:"required" `
	Password string `json:"password" uri:"password" yaml:"password" form:"password" binding:"required" `
}

func main() {

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {

		var u User
		err := c.BindJSON(&u)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "ERROR",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "OK",
				"data": u,
			})
		}
	})

	err := r.Run(":8001")
	if err != nil {
		return
	}

}
