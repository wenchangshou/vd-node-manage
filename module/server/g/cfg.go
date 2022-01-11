package g

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/toolkits/file"
)

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}
type DatabaseConfig struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	TablePrefix string `json:"tablePrefix"`
	DBFile      string `json:"dbFile"`
	Port        int    `json:"port"`
}
type CacheConfig struct {
	Passwd     string   `json:"passwd"`
	DB         int      `json:"db"`
	Addr       string   `json:"addr"`
	MaxIdle    int      `json:"maxIdle"`
	HighQueues []string `json:"highQueues"`
	LowQueues  []string `json:"lowQueues"`
}
type RedisConfig struct {
	Passwd     string   `json:"passwd"`
	DB         int      `json:"db"`
	Addr       string   `json:"addr"`
	MaxIdle    int      `json:"maxIdle"`
	HighQueues []string `json:"highQueues"`
	LowQueues  []string `json:"lowQueues"`
}
type GlobalConfig struct {
	Debug    bool            `json:"debug"`
	Mode     string          `json:"mode"`
	Hosts    string          `json:"hosts"`
	MaxCons  int             `json:"maxCons"`
	MaxIdle  int             `json:"maxIdle"`
	Listen   string          `json:"listen"`
	Http     *HttpConfig     `json:"http"`
	Database *DatabaseConfig `json:"database"`
	Cache    *CacheConfig    `json:"cache"`
	Redis    *RedisConfig    `json:"redis"`
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
		log.Fatalln("config file:", cfg, "is not existent")
	}
	ConfigFile = cfg
	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file,", cfg, "fail:", err)
	}
	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}
	configLock.Lock()
	defer configLock.Unlock()
	config = &c
	log.Println("read config file:", cfg, "successfully")
}
