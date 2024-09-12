package v1

import (
	"fmt"
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

func UserLogin(c *gin.Context) {

	fmt.Println(c.Request.Context())
	fmt.Println(c.Request.Header)
	fmt.Println(c.Request.Body)
	fmt.Println(c.Request.Cookies())
	fmt.Println(c.Request.Host)
	fmt.Println(c.Request.Method)

	var userLoginService service.UserService

	if err := c.ShouldBind(&userLoginService); err == nil {

		res := userLoginService.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
