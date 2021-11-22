package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync"
)


// SystemConfig 系统配置
type SystemConfig struct {
	Init  bool   `json:"init"`
	Mode  string `json:"mode"`
	IP    string `json:"ip"`
	Port  uint   `json:"port"`
	Debug bool   `json:"debug"`
}
// ServerConfig 服务配置
type ServerConfig struct {
	Address string `json:"address"`
}
// LogConfig 日志配置
type LogConfig struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Ext        string `json:"ext"`
	Level      string `json:"level"`
	MaxMsgSize int    `json:"maxMsgSize"`
}
// RpcConfig rpc配置
type RpcConfig struct {
	Address string `json:"address"`
}
// ResourceConfig 资源配置
type ResourceConfig struct {
	Directory string `json:"directory"`
	Tmp       string `json:"tmp"`
}
type TaskConfig struct {
	Count int `json:"count"`
}
type GlobalConfig struct {
	Debug    bool           `json:"debug"`
	Log      LogConfig      `json:"log"`
	System   SystemConfig   `json:"system"`
	Server   ServerConfig   `json:"server"`
	Rpc      RpcConfig      `json:"rpc"`
	Resource ResourceConfig `json:"resource"`
	Task     TaskConfig     `json:"task"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)
// Config 返回配置
func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// ParseConfig 解析配置文件
func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("config file:", cfg, "is not exists")
	}
	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not exists")
	}
	ConfigFile = cfg
	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}
	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file", cfg, "error:", err.Error())
	}
	configLock.Lock()
	defer configLock.Unlock()
	config = &c
	log.Println("g.ParseConfig ok,file", cfg)
}
