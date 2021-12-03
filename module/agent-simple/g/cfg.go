package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"os"
	"sync"
)

// ServerConfig 服务配置
type ServerConfig struct {
	Address           string `json:"address"`
	Mode              string `json:"mode"`
	HttpAddress       string `json:"httpAddress"`
	RpcAddress        string `json:"rpcAddress"`
	Register          bool   `json:"register"`
	ID                uint   `json:"id"`
	ReportInterval    int    `json:"reportInterval"`
	QueryTaskInterval int    `json:"queryTaskInterval"`
	RedisAddress      string `json:"redis_address"`
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
type HttpConfig struct {
	Enabled  bool   `json:"enabled"`
	Listen   string `json:"listen"`
	Backdoor bool   `json:"backdoor"`
}
type HeartbeatConfig struct {
	Enabled  bool   `json:"enabled"`
	Addr     string `json:"addr"`
	Interval int    `json:"interval"`
	Timeout  int    `json:"timeout"`
}
type GlobalConfig struct {
	Debug     bool             `json:"debug"`
	Hostname  string           `json:"hostname"`
	IP        string           `json:"ip"`
	Log       *LogConfig       `json:"log"`
	Server    *ServerConfig    `json:"server"`
	Rpc       *RpcConfig       `json:"rpc-old"`
	Resource  *ResourceConfig  `json:"resource"`
	Task      *TaskConfig      `json:"task"`
	Heartbeat *HeartbeatConfig `json:"heartbeat"`
	Http      *HttpConfig      `json:"http"`
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
func SetRegisterStatus(status bool, id uint, httpAddress, rpcAddress, redisAddress string) {
	configLock.Lock()
	defer configLock.Unlock()
	config.Server.Register = status
	config.Server.ID = id
	config.Server.HttpAddress = httpAddress
	config.Server.RpcAddress = rpcAddress
	config.Server.RedisAddress = redisAddress
	b, _ := json.Marshal(config)
	file.WriteString(ConfigFile, string(b))
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

func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != "" {
		return hostname, nil
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("ERROR: os.Hostname() fail", err)
	}
	return hostname, err
}
func IP() string {
	ip := Config().IP
	if ip != "" {
		return ip
	}
	if len(LocalIp) > 0 {
		ip = LocalIp
	}
	return ip
}
