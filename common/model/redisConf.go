package model

type RedisConfig struct {
	Passwd     string   `json:"passwd"`
	DB         int      `json:"db"`
	Addr       string   `json:"addr"`
	MaxIdle    int      `json:"maxIdle"`
	HighQueues []string `json:"highQueues"`
	LowQueues  []string `json:"lowQueues"`
}
