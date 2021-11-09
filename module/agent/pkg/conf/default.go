package conf

// DatabaseConfig 数据库配置
var DatabaseConfig = &database{
	Type:   "UNSET",
	DBFile: "vd.db",
	Port:   3306,
}
var SystemConfig = &system{
	Debug: false,
	Mode:  "auto",
}
var LogConfig = &log{
	Name:         "log",
	Path:         "logs",
	Ext:          "log",
	Level:        "debug",
	MaxMsgSize:   1048576,
	ArgumentType: "file",
}

var RpcConfig = &rpc{
	Address: "localhost:10051",
}
var ResourceConfig = &Resource{
	Directory: "c:/zoolon",
	Tmp:       "./tmp",
}

var ServerConfig = &Server{
	Address: "0.0.0.0:8888",
}
var TaskConfig=&Task{
	Count: 1,
}