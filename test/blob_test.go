package test

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"os"
	"testing"
	"time"
)

var testBucket = []byte("test-bucket")

func TestBlob(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		db.Close()
		os.RemoveAll("my.db")
	}()
	key := []byte("hello")
	value := []byte("world")
	// 创建一个 read-write 事务来进行写操作
	err = db.Update(func(tx *bolt.Tx) error {
		// 如果 bucket 不存在则，创建一个 bucket
		bucket, err := tx.CreateBucketIfNotExists(testBucket)
		if err != nil {
			return err
		}

		// 将 key-value 写入到 bucket 中
		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	var k2, v2 string
	// 创建一个 read-only 事务来获取数据
	err = db.View(func(tx *bolt.Tx) error {
		// 获取对应的 bucket
		bucket := tx.Bucket(testBucket)
		// 如果 bucket 返回为 nil，则说明不存在对应 bucket
		if bucket == nil {
			return fmt.Errorf("Bucket %q is not found", testBucket)
		}
		// 从 bucket 中获取对应的 key（即上面写入的 key-value）
		val := bucket.Get(key)
		fmt.Printf("%s: %s\n", string(key), string(val))
		k2 = string(key)
		v2 = string(val)
		return nil
	})
	fmt.Println("k2,v2", k2, v2)
	if err != nil {
		log.Fatal(err)
	}
}
