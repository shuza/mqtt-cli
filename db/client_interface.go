package db

/**
 * :=  created by:  Shuza
 * :=  create date:  01-May-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

type IDbClient interface {
	Init()
	Set(key string, value string) error
	Get(key string) string
	Close()
}

var Client IDbClient
