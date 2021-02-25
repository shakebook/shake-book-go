package dao

import (
	"context"
	"log"
	"time"
)

//WriteEmailCodeToRedis implement Conn
func (c *Conn) WriteEmailCodeToRedis(ctx context.Context, k, v string) error {
	rdb := c.Redis.RedisClient()
	log.Printf("WriteEmailCodeToRedis:k:%s,v:%s\n", k, v)
	return rdb.Set(ctx, k, v, 1*time.Minute).Err()
}
