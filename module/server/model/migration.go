package model

import (
	"fmt"
	"log"
)

func needMigration() bool {
	return true
}
func migration() error {
	var (
		err error
	)
	// config := g.Config()
	if !needMigration() {
		return nil
	}
	// if config.Database.Type == "mysql" {
	//DB = DB.Set("gorm:table_options", "ENGINE=InnoDB")
	// }
	err = DB.AutoMigrate(
		&Device{},
		&DeviceResource{},
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
