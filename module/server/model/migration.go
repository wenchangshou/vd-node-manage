package model

import (
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/server/g"
	"log"
)

func needMigration() bool {
	return true
}
func migration() error {
	var (
		err error
	)
	config := g.Config()
	if !needMigration() {
		return nil
	}
	if config.Database.Type == "mysql" {
		//DB = DB.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	err = DB.AutoMigrate(
		&Device{},
		&File{},
		&Resource{},
		Task{},
		Event{},
		&ResourceDistribution{},
	)
	if err != nil {
		log.Fatalln("migration database fail", "error", err)
	}
	fmt.Println("数据库迁移结束")
	return nil
}
