package g

import (
	"encoding/json"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/agent-simple/g/db"
	bolt "go.etcd.io/bbolt"
	"sync"
)

var (
	serverDBLock *sync.RWMutex
	serverInfo   *model.ServerConfig
)

func init() {
	serverDBLock = new(sync.RWMutex)
	serverInfo = &model.ServerConfig{Register: false}
}

func GetServerInfo() model.ServerConfig {
	serverDBLock.RLock()
	defer serverDBLock.RUnlock()
	return *serverInfo
}
func ResetServerInfo() {
	serverDBLock.Lock()
	serverInfo = &model.ServerConfig{Register: false}
	serverDBLock.Unlock()
	StoreServerInfo(serverInfo)
}
func StoreServerInfo(info *model.ServerConfig) error {
	serverDBLock.Lock()
	defer serverDBLock.Unlock()
	serverInfo = info
	err := db.GDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("config"))
		if err != nil {
			return err
		}
		b, err := json.Marshal(info)
		if err != nil {
			return err
		}
		return bucket.Put([]byte("server"), b)
	})
	return err
}
func SetSettings(bucket string, values map[string]string, prefix string) error {
	err := db.GDB.Batch(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		for k, v := range values {
			b.Put([]byte(prefix+k), []byte(v))
		}
		return nil
	})
	return err
}
func GetSettings(bucket string, keys []string, prefix string) map[string]string {
	rtu := make(map[string]string)
	db.GDB.Batch(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		for _, k := range keys {
			v := b.Get([]byte(prefix + k))
			rtu[k] = string(v)
		}
		return nil
	})
	return rtu
}
func LoadServerInfoByDb() error {
	serverDBLock.RLock()
	var info *model.ServerConfig
	err := db.GDB.View(func(tx *bolt.Tx) error {
		_server := model.ServerConfig{}
		bucket := tx.Bucket([]byte("config"))
		if bucket == nil {
			return nil
		}
		b := bucket.Get([]byte("server"))
		err := json.Unmarshal(b, &_server)
		if err != nil {
			return err
		}
		info = &_server
		return nil
	})
	serverDBLock.RUnlock()
	if err != nil {
		return err
	}
	if info == nil {
		info = &model.ServerConfig{
			Register: false,
		}
		StoreServerInfo(info)
	}
	serverInfo = info
	return nil

}
