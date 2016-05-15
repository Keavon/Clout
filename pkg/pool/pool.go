package pool

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// NewPool creates a redis pool
func NewPool(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(url)
		},
		TestOnBorrow: func(rc redis.Conn, t time.Time) error {
			_, err := rc.Do("PING")
			return err
		},
	}
}
