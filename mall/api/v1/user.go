package v1

import (
	"github.com/gin-gonic/gin"
	"mall/service"
	"net/http"
)

func UserRegister(c *gin.Context) {

	// 相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
