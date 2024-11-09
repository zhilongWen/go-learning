package init_mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"test-server/config"
)

var DB *gorm.DB

func InitMySQL(admin config.Admin) *gorm.DB {
	db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@tcp("+admin.Path+")/"+admin.DbName+"?"+admin.Config)
	if err != nil {
		log.Println("Failed to connect to MySQL database")
	} else {
		DB = db
		DB.DB().SetMaxOpenConns(10)
		DB.DB().SetMaxIdleConns(5)
	}
	return DB
}
