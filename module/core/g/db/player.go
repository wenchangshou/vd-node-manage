package db

import (
	"encoding/json"
	"github.com/wenchangshou/vd-node-manage/common/model"
	bolt "go.etcd.io/bbolt"
)

func AddPlayer(service string, player *model.Player) error {

	err := GDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("player"))
		bucket.Put([]byte(service), player.Serialization())
		return nil
	})
	return err
}
func GetPlayer(service string) (player *model.Player) {
	GDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("player"))
		b := bucket.Get([]byte(service))
		return json.Unmarshal(b, &player)
	})
	return
}
func ListPlayer() (players []*model.Player) {
	GDB.View(func(tx *bolt.Tx) error {
		buc
	})
}
func GetSetting(conf []byte) []byte {
	var (
		rtu []byte
	)
	GDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("config"))
		rtu = bucket.Get(conf)
		return err
	})
	return rtu
}
func SetSetting(conf []byte, val []byte) error {
	return GDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("config"))
		if err != nil {
			return err
		}
		err = bucket.Put(conf, val)
		return err
	})
}
