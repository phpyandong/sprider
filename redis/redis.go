package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisClient *redis.Pool

// 初始化
func InitRedis() {
	RedisClient = &redis.Pool{
		MaxIdle:     4,
		MaxActive:   4,
		IdleTimeout: time.Duration(2 * time.Second),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "140.143.139.224:6379",
				redis.DialConnectTimeout(5*time.Second),
				redis.DialReadTimeout(3*time.Second),
				redis.DialWriteTimeout(3*time.Second))
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", "DeNpDgQaXPR92xDrw"); err != nil {
				c.Close()
				return nil, err
			}

			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func GetRedis() *redis.Pool {
	return RedisClient
}