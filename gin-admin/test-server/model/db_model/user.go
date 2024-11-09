package db_model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"test-server/init/init_mysql"
)

type User struct {
	gorm.Model `json:"-"`
	UUID       uuid.UUID `json:"uuid"`
	UserName   string    `json:"userName"`
	PassWord   string    `json:"passWord"`
	NickName   string    `json:"nickName" gorm:"default:'galeone'"`
	HeaderImg  string    `json:"headerImg" gorm:"default:'galeone'"`
}

func (u *User) Create() (err error, user interface{}) {
	err = init_mysql.DB.Create(u).Error
	return err, user
}
