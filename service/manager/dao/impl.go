package dao

import (
	"database/sql"

	redis "github.com/go-redis/redis/v8"
)

//Conn implement
type Conn struct {
	Mysql MysqlConn
	Redis RedisConn
}

//MysqlConn db
type MysqlConn interface {
	DB() *sql.DB
}

//RedisConn client
type RedisConn interface {
	RedisClient() *redis.Client
}
