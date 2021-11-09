package model

import "strings"

type Module struct {
	Base
	Category string `json:"category" form:"category"`
	Name     string `gorm:"name" json:"name"`
	Value    string `gorm:"value" json:"value"`
}

func (module *Module) TableName() string {
	return "module"
}
func (module *Module) Create() error {
	return DB.Model(&Module{}).Create(&module).Error
}

func DeleteModuleByID(id string) error {
	return DB.Unscoped().Delete(&Module{}, id).Error
}

//ListModules 获取模块列表
func ListModules(page int, size int, orderBy string, conditions map[string]string, searches map[string]string) ([]Module, int64) {
	var res []Module
	var total int64
	tx := DB.Model(&Module{})
	if orderBy != "" {
		tx = tx.Order(orderBy)
	}
	for k, v := range conditions {
		tx = tx.Where(k+" = ?", v)
	}
	if len(searches) > 0 {
		search := ""
		for k, v := range searches {
			search += (k + " like '%" + v + "%' OR ")
		}
		search = strings.TrimSuffix(search, " OR ")
		tx = tx.Where(search)
	}
	tx.Count(&total)
	tx.Debug().Limit(size).Offset((page - 1) * size).Find(&res)
	return res, total
}
