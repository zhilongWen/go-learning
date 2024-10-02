package routes

import (
	"github.com/gin-gonic/gin"
	api "mall/api/v1"
	"mall/middleware"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.Use(middleware.Cors())
	//router.StaticFile("/static", string(http.Dir("/static")))
	router.StaticFile("/static", "./static")
	v1 := router.Group("api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		// 用户操作
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("/user/login", api.UserLogin)
		// /static/imgs/avatar
		// /static/imgs/avatar/avatar.JPG

		// 用户操作需要登陆保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar) // 上传头像
		}

	}

	return router
}
