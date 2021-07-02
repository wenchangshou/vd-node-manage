package conf

// DatabaseConfig 数据库配置
var DatabaseConfig = &database{
	Type:   "UNSET",
	DBFile: "vd.db",
	Port:   3306,
}
var ServerConfig = &server{
	Ip:   "127.0.0.1",
	Port: 1111,
}

// CORSConfig 跨域配置
var CORSConfig = &cors{
	AllowOrigins:     []string{"UNSET"},
	AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS", "DELETE", "PATCH"},
	AllowHeaders:     []string{"Cookie", "X-Policy", "Authorization", "Content-Length", "Content-Type", "X-Path", "X-FileName"},
	AllowCredentials: true,
	ExposeHeaders:    nil,
}
var SystemConfig = &system{
	Debug:  false,
	Mode:   "single",
	Listen: ":8888",
}
var LogConfig = &log{
	Name:         "log",
	Path:         "logs",
	Ext:          "log",
	Level:        "debug",
	MaxMsgSize:   1048576,
	ArgumentType: "file",
}
