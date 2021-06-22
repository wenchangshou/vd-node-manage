package conf

import (
	"errors"

	"github.com/go-ini/ini"
	"github.com/go-playground/validator/v10"
	"github.com/wenchangshou2/vd-node-manage/pkg/util"
	"github.com/wenchangshou2/zutil"
)

//system 系统通用配置
type system struct {
	Mode          string `validate:"eq=standard|eq=mixing"`
	Listen        string `validate:"required"`
	Debug         bool
	SessionSecret string
	HashIDSalt    string
}

// database 数据库
type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	DBFile      string
	Port        int
}
type cors struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
}

type log struct {
	Name         string
	Path         string
	Ext          string
	Level        string
	MaxMsgSize   int
	ArgumentType string
}

var cfg *ini.File

const defaultConf = `
[System]
Listen = ":8888"
SessionSecret = {SessionSecret}
HashIDSalt = {HashIDSalt}
`

func Init(path string) error {
	var err error
	if path == "" || !zutil.IsExist(path) {
		// 创建初始配置文件
		confContent := util.Replace(map[string]string{
			"{SessionSecret}": zutil.RandStringRunes(64),
			"{HashIDSalt}":    zutil.RandStringRunes(64),
		}, defaultConf)
		f, err := zutil.CreatNestedFile(path)
		if err != nil {
			panic("无法创建配置文件," + err.Error())
		}
		_, err = f.WriteString(confContent)
		if err != nil {
			return errors.New("无法配置文件")
		}
		f.Close()
	}
	cfg, err = ini.Load(path)
	if err != nil {
		return errors.New("无法解析配置文件:" + path + " : " + err.Error())
	}
	sections := map[string]interface{}{
		"Database": DatabaseConfig,
		"System":   SystemConfig,
		"CORS":     CORSConfig,
	}
	for sectionName, sectionStruct := range sections {
		err = mapSection(sectionName, sectionStruct)
		if err != nil {
			return errors.New("配置文件 " + sectionName + ",解析失败:" + err.Error())
		}
	}
	return nil
}

// mapSection 将配置文件的 Section 映射到结构体上
func mapSection(section string, confStruct interface{}) error {
	err := cfg.Section(section).MapTo(confStruct)
	if err != nil {
		return err
	}

	// 验证合法性
	validate := validator.New()
	err = validate.Struct(confStruct)
	if err != nil {
		return err
	}

	return nil
}
