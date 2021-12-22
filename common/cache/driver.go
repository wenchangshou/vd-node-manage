package cache

import (
	"errors"
	"strconv"
)

// Store 缓存存储器
//var Store Driver = NewMemoStore()
var Store Driver

// InitCache 创建新的
func InitCache(provider string, address string, password string, db int) error {
	if provider == "mem" {
		Store = NewMemoStore()
		return nil
	} else if provider == "redis" {
		s := strconv.Itoa(db)
		Store = NewRedisStore(10, "tcp", address, password, s)
		return nil
	}
	return errors.New("未找到相应的cache提供者")

}

type Driver interface {
	// Set 设置一个key
	Set(key string, value interface{}, ttl int) error

	Get(key string) (interface{}, bool)
	Gets(keys []string, prefix string) (map[string]interface{}, []string)
	// Sets 批量设置值，所有的key都会加上prefix前缀
	Sets(values map[string]interface{}, prefix string) error
	Delete(keys []string, prefix string) error
}

func Set(key string, value interface{}, ttl int) error {
	return Store.Set(key, value, ttl)
}
func Get(key string) (interface{}, bool) {
	return Store.Get(key)
}
func Deletes(keys []string, prefix string) error {
	return Store.Delete(keys, prefix)
}
func GetSettings(keys []string, prefix string) (map[string]string, []string) {
	raw, miss := Store.Gets(keys, prefix)
	res := make(map[string]string, len(raw))
	for k, v := range raw {
		res[k] = v.(string)
	}
	return res, miss
}
func SetSettings(values map[string]string, prefix string) error {
	var toBeSet = make(map[string]interface{}, len(values))
	for key, value := range values {
		toBeSet[key] = interface{}(value)
	}
	return Store.Sets(toBeSet, prefix)
}
