package g

import (
	"log"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConnPool *redis.Pool

func InitRedisConnPool() {
	redisConfig := Config().Redis
	auth, addr := formatRedisAddr(redisConfig.Addr)
	RedisConnPool = &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: pingRedis,
	}

}
func formatRedisAddr(addrConfig string) (string, string) {
	if redisAddr := strings.Split(addrConfig, "@"); len(redisAddr) == 1 {
		return "", redisAddr[0]
	} else {
		return strings.Join(redisAddr[0:len(redisAddr)-1], "@"), redisAddr[len(redisAddr)-1]
	}
}
func pingRedis(c redis.Conn, _ time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
	}
	return err
}
