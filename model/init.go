package model

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	logging.G_Logger.Info("初始化数据库连接")
	var (
		db  *gorm.DB
		err error
	)
	if gin.Mode() == gin.TestMode {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		switch conf.DatabaseConfig.Type {
		case "UNSET", "sqlite", "sqlite3":
			db, err = gorm.Open(sqlite.Open(conf.DatabaseConfig.DBFile), &gorm.Config{})
		case "mysql":
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.Port,
				conf.DatabaseConfig.Name)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		default:
			panic("不支持数据库类型:" + conf.DatabaseConfig.Type)
		}
	}
	if err != nil {
		panic("连接数据库不成功," + err.Error())
	}
	// _db, err := db.DB()
	// _db.SetMaxIdleConns(50)
	// _db.SetMaxOpenConns(100)
	// _db.SetConnMaxLifetime(time.Second * 30)
	DB = db
	err = migration()
	if err != nil {
		return err
	}
	return nil
}
