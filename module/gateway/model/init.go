package model

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/g"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	var (
		db  *gorm.DB
		err error
	)
	cfg:=g.Config()
	logging.GLogger.Info("初始化数据库连接")
	if gin.Mode() == gin.TestMode {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {

		switch cfg.Database.Type {
		case "UNSET", "sqlite", "sqlite3":
			db, err = gorm.Open(sqlite.Open(cfg.Database.DBFile), &gorm.Config{})
		case "mysql":
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cfg.Database.User,
				cfg.Database.Password,
				cfg.Database.Host,
				cfg.Database.Port,
				cfg.Database.Name)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		default:
			panic("不支持数据库类型:" + cfg.Database.Type)
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
