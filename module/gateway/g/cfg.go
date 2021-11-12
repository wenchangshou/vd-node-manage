package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type SystemConfig struct {
	Mode          string `json:"mode"`
	Listen        string `validate:"required" json:"listen"`
	Debug         bool   `json:"debug"`
	SessionSecret string `json:"sessionSecret"`
	HashIDSalt    string `json:"hashIDSalt"`
	IP string `json:"ip"`
	Port uint `json:"port"`
}
type DatabaseConfig struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	TablePrefix string `json:"tablePrefix"`
	DBFile      string `json:"DBFile"`
	Port        int    `json:"port"`
}
type CorsConfig struct {
	AllowOrigins     []string `json:"allowOrigins"`
	AllowMethods     []string `json:"allowMethods"`
	AllowHeaders     []string `json:"allowHeaders"`
	AllowCredentials bool     `json:"allowCredentials"`
	ExposeHeaders    []string `json:"exposeHeaders"`
}
type ServerConfig struct {
	IP   string `json:"ip"`
	Port uint   `json:"port"`
}
type ZebusConfig struct {
	IP       string `json:"ip"`
	HttpPort int    `json:"httpPort"`
	WsPort   int    `json:"WsPort"`
}
type LogConfig struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	Ext          string `json:"ext"`
	Level        string `json:"level"`
	MaxMsgSize   int    `json:"maxMsgSize"`
	ArgumentType string `json:"argumentType"`
}
type GlobalConfig struct {
	System   *SystemConfig   `json:"system"`
	Database *DatabaseConfig `json:"database"`
	Cors     *CorsConfig     `json:"cors"`
	Server   *ServerConfig   `json:"server"`
	Zebus    *ZebusConfig    `json:"zebus"`
	Log      LogConfig       `json:"log"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
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
