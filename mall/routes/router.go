package routes

import (
	"github.com/gin-gonic/gin"
	api "mall/api/v1"
	"mall/middleware"
	"net/http"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.Use(middleware.Cors())
	router.StaticFile("/static", string(http.Dir("/static")))
	v1 := router.Group("api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		// 用户操作
		v1.POST("user/register", api.UserRegister)

	}

	return router
}
