package model

type ServerRedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
type ServerHttpConfig struct {
	Enable  bool   `json:"enable"`
	Address string `json:"address"`
}
type ServerRpcConfig struct {
	Enable  bool   `json:"enable"`
	Address string `json:"address"`
}
type ServerEventConfig struct {
	Provider  string            `json:"provider"`
	Arguments map[string]string `json:"arguments"`
}
type ServerConfig struct {
	Register bool              `json:"register"`
	Server   string            `json:"server"`
	ID       uint              `json:"id"`
	Redis    ServerRedisConfig `json:"redis"`
	Http     ServerHttpConfig  `json:"http"`
	Rpc      ServerRpcConfig   `json:"rpc"`
}
