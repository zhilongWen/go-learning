package dao

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

func Database(connReadPath, connWritePath string) {

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN:                       connReadPath, // DSN data source name
				DefaultStringSize:         256,          // string 类型字段的默认长度
				DisableDatetimePrecision:  true,         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
				DontSupportRenameIndex:    true,         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
				DontSupportRenameColumn:   true,         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
				SkipInitializeWithVersion: false,        // 根据版本自动配置
			},
		),
		&gorm.Config{
			Logger: ormLogger,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)

	if err != nil {
		fmt.Println("Error connection database:", err)
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	_db = db

	// 读写分离，主从配置
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			// `db2` 作为 sources，`db3`、`db4` 作为 replicas
			Sources:  []gorm.Dialector{mysql.Open(connReadPath)},                            // 写操作
			Replicas: []gorm.Dialector{mysql.Open(connWritePath), mysql.Open(connReadPath)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                             // sources/replicas 负载均衡策略
		}))

	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
