package db

import (
	bolt "go.etcd.io/bbolt"
)

// BoltDb Bolt存储库
type BoltDb struct {
	db *bolt.DB
}

func NewBoltDb(dbPath string) (*BoltDb, error) {
	db, err := bolt.Open(dbPath, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &BoltDb{db}, nil
}
