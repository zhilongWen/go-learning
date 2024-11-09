package register_table

import (
	"github.com/jinzhu/gorm"
	"test-server/model/db_model"
)

func RegistTable(db *gorm.DB) {
	db.AutoMigrate(db_model.User{})
}
