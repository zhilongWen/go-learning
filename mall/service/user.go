package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	util "mall/pkg/utils"
	"mall/serializer"
)

// UserService 管理用户服务
type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

func (service UserService) Register(ctx context.Context) serializer.Response {

	var user model.User
	code := e.SUCCESS

	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥长度不足",
		}
	}

	// key 加密
	util.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)

	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{Status: code, Msg: e.GetMsg(code)}
	}

	if exist {
		code = e.ErrorExistUser
		return serializer.Response{Status: code, Msg: e.GetMsg(code)}
	}

	user = model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额
	}

	// 加密密码
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登陆函数
func (service UserService) Login(ctx context.Context) serializer.Response {

	var user *model.User

	code := e.SUCCESS

	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)

	if !exist {
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}
