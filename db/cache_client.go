package db

import (
	"github.com/boltdb/bolt"
	"os"
	"strings"
	"time"
)

/**
 * :=  created by:  Shuza
 * :=  create date:  01-May-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

type CacheClient struct {
	db *bolt.DB
}

func (c *CacheClient) Init() {
	path := getDbPath()
	db, err := bolt.Open(path+"/mqtt-sh.db", 0777, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		panic(err)
	}
	c.db = db
	c.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte(Bucket))
		return nil
	})
}

func (c *CacheClient) Set(key string, value string) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(Bucket))
		err := bucket.Put([]byte(key), []byte(value))
		return err
	})
}

func (c *CacheClient) Get(key string) string {
	value := ""
	c.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(Bucket))
		v := bucket.Get([]byte(key))
		value = string(v)
		return nil
	})
	return value
}

func (c *CacheClient) Close() {
	c.db.Close()
}

func getDbPath() string {
	cacheDir, _ := os.UserCacheDir()
	parts := strings.Split(cacheDir, ".")
	path := parts[0] + "Downloads/.mqtt-sh/tmp/"
	os.MkdirAll(path, 0777)
	return path
}
