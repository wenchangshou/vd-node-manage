package model

import (
	"github.com/wenchangshou2/vd-node-manage/common/logging"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/g"
)

func needMigration() bool {
	return true
}
func migration() error {
	var (
		err error
	)
	if !needMigration() {
		logging.GLogger.Info("数据库版本匹配,跳过数据库迁移")
		return nil
	}
	logging.GLogger.Info("开始进行数据库初始化......")
	if g.Config().Database.Type == "mysql" {
		DB = DB.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	err = DB.AutoMigrate(&User{}, &Setting{}, &Project{}, &File{}, &ProjectRelease{},
		&Computer{},
		&Resource{},
		&Task{},
		&TaskItem{},
		&ExhibitionCategory{},
		&Exhibition{},
		&CustomLayout{},
		&Window{},
		&Module{},
		&ExhibitionWindowItem{},
	)
	if err != nil {
		logging.GLogger.Error("数据库迁移失败:" + err.Error())
		return err
	}
	//创建初始存储策略
	logging.GLogger.Info("数据库初始化结束")
	return err
}
func addDefaultSettings() {
	defaultSettings := []Setting{
		{
			Name: "siteURL", Value: `http://localhost`, Type: "basic",
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
