package init_router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"test-server/middleware"
	"test-server/router"
)

var Router = gin.Default()

func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.Use(middleware.Logger())
	router.InitUserRouter(Router)
	return Router
}
