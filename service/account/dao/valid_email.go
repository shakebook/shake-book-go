package dao

import (
	"context"
	"errors"
	"log"
)

//ValidEmail implement Conn
func (c *Conn) ValidEmail(ctx context.Context, email, reqCode string) error {
	rdb := c.Redis.RedisClient()
	val, err := rdb.Get(ctx, email).Result()
	log.Printf("ValidEmail:val:%s,err:%v\n", val, err)
	if err != nil {
		return errors.New("验证码已过期")
	}
	if val != reqCode {
		return errors.New("验证码不正确")
	}
	return nil
}
