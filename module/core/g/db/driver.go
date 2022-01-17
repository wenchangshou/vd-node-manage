package db

import (
	bolt "go.etcd.io/bbolt"
)

var Store Driver

type Driver interface {
	GetResource(id uint) (map[string]string, bool)
	SetResource(id uint, val map[string]string) error
}

var GDB *bolt.DB

func InitDB(dbpath string) error {
	var (
		err error
	)
	GDB, err = bolt.Open(dbpath, 0666, nil)
	return err
}
