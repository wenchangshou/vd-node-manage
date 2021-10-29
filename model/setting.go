package model

import (
	"net/url"
	"strconv"
)

type Setting struct {
	Base
	Type  string `gorm:"not null"`
	Name  string `gorm:"unique;not null;index:setting_key"`
	Value string `gorm:"size:‎65535"`
}

// IsTrueVal 返回设置的值是否为真
func IsTrueVal(val string) bool {
	return val == "1" || val == "true"
}

// GetSettingByName 用 Name 获取设置值
func GetSettingByName(name string) string {
	var setting Setting

	// 尝试数据库中查找
	result := DB.Where("name = ?", name).First(&setting)
	if result.Error == nil {
		return setting.Value
	}
	return ""
}
func GetSiteURL() *url.URL {
	base, err := url.Parse(GetSettingByName("siteURL"))
	if err != nil {
		base, _ = url.Parse("http://localhost")
	}
	return base
}

// GetSettingByNames 用多个 Name 获取设置值
func GetSettingByNames(names ...string) map[string]string {
	return nil
}

// GetIntSetting 获取整形设置值，如果转换失败则返回默认值defaultVal
func GetIntSetting(key string, defaultVal int) int {
	res, err := strconv.Atoi(GetSettingByName(key))
	if err != nil {
		return defaultVal
	}
	return res
}
