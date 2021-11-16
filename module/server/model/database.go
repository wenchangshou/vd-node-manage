package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/module/server/g"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	var (
		db  *gorm.DB
		err error
	)
	config := g.Config()
	if gin.Mode() == gin.TestMode {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		switch config.Database.Type {
		case "UNSET", "sqlite", "sqlite3":
			db, err = gorm.Open(sqlite.Open(config.Database.DBFile), &gorm.Config{})
		case "mysql":
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
				config.Database.User,
				config.Database.Password,
				config.Database.Host,
				config.Database.Port,
				config.Database.Name)
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		default:
			log.Fatalln("Unsupported database type")
		}
	}
	if err != nil {
		log.Fatalln("Unable to connect to the database", "err", err)
	}
	DB=db
	err=
}
