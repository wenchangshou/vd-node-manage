package model

import (
	"github.com/wenchangshou2/vd-node-manage/pkg/conf"
	"github.com/wenchangshou2/vd-node-manage/pkg/logging"
)

func needMigration() bool {
	return true
}
func migration() error {
	var (
		err error
	)
	if !needMigration() {
		logging.G_Logger.Info("数据库版本匹配,跳过数据库迁移")
		return nil
	}
	logging.G_Logger.Info("开始进行数据库初始化......")
	if conf.DatabaseConfig.Type == "mysql" {
		DB = DB.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	err = DB.AutoMigrate(&User{}, &Setting{}, &Project{}, &File{}, &ProjectRelease{}, &Computer{}, &Resource{}, &Task{}, &TaskItem{}, &ComputerProject{}, &ComputerResource{})
	if err != nil {
		logging.G_Logger.Error("数据库迁移失败:" + err.Error())
		return err
	}
	//创建初始存储策略
	logging.G_Logger.Info("数据库初始化结束")
	return err
}
func addDefaultSettings() {
	defaultSettings := []Setting{
		{
			Name: "siteURL", Value: `http://localhost`, Type: "biasic",
		},
		{
			Name: "siteName", Value: `vd-node-manage`, Type: "basic",
		},
		{
			Name: "runModel", Value: "master", Type: "system",
		},
	}
	for _, value := range defaultSettings {
		DB.Where(Setting{Name: value.Name}).Create(&value)
	}
}
