package cron_server

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

func pool() *redis.Pool {
	host := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	pool := &redis.Pool{
		IdleTimeout: 240 * time.Second, // close connections after remaining idle for this duration
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host, redis.DialDatabase(1))
			if err != nil {
				return nil, err
			}

			//if _, err := c.Do("AUTH", i.redisConfig.password); err != nil {
			//	log.Println("redigo-wrapper: Redis: AUTH failed", err)
			//	_ = c.Close()
			//	return nil, err
			//}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return pool
}
