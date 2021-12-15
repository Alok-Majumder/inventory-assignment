package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisPoolConn struct {
	redisPoolConn *redis.Pool
}

//Create Connection with reids instance
func NewRedisPollConn(redisHost string, redisPort string) (*RedisPoolConn, error) {

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	pool := &redis.Pool{
		MaxIdle:         80,
		MaxActive:       160,
		MaxConnLifetime: 0,
		IdleTimeout:     60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddr)
			if err != nil {
				return nil, fmt.Errorf("redis.Dial: %v", err)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			fmt.Println("Error on PING ")

			return err
		},
	}
	redispool := &RedisPoolConn{
		pool,
	}
	fmt.Println("Connection Created")

	return redispool, nil
}

//Set multiple keys-values in redis hash
func (r *RedisPoolConn) HMSet(ctx context.Context, hkey string, keyValueMap map[string]interface{}) error {
	conn := r.redisPoolConn.Get()
	defer conn.Close()

	_, err := conn.Do("HMSET", redis.Args{}.Add(hkey).AddFlat(keyValueMap)...)
	if err != nil {
		return fmt.Errorf("redis HMSET failed for data: %v due to %q", keyValueMap, err.Error())

	}

	return nil

}
