package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test-server/model/db_model"
)

type RegistStuct struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

// Regist
// @Summary 创建用户
// @Produce  application/x-www-form-urlencoded
// @Param data body api.RegistStuct true "用户注册接口"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/regist [post]
func Regist(c *gin.Context) {
	var U db_model.User
	_ = c.BindJSON(&U)
	fmt.Println(U)
	err, user := U.Create()
	fmt.Println(err, user)
}
