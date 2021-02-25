package conn

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

//Redis 实现者
type Redis struct {
	Addr     string
	Password string
	Context  context.Context
}

//RedisClient redis client
func (r *Redis) RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       0, // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("redis connect failed, error:%v\n", err.Error())
	}
	fmt.Println(pong, err)
	return rdb
}
