package redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"time"
)

var poll *redis.Pool

// Setup Initialize the Redis instance
func Setup(config *Config) error {
	config.FillWithDefaults()
	poll = &redis.Pool{
		Dial:        config.Connect,                                // 提供创建和配置应用程序连接
		MaxIdle:     config.MaxIdle,                                // 最大空闲连接数
		MaxActive:   config.MaxActive,                              // 最大活动连接数
		IdleTimeout: time.Duration(config.Timeout*2) * time.Second, // 空闲连接超时时间
		Wait:        true,                                          // 如果Wait被设置成true，则Get()方法将会阻塞
	}
	return ping()
}

func ping() error {
	conn := poll.Get()
	defer conn.Close()
	_, err := conn.Do("ping")
	return err
}

// Get all key
func GetAllKeys() []string {
	conn := poll.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		return make([]string, 0)
	}
	return keys
}

// check key exists
func Exists(key string) bool {
	conn := poll.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get key
func Get(key string) (string, error) {
	conn := poll.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

// Delete key
func Delete(key string) (bool, error) {
	conn := poll.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// Set key/value
func Set(key string, data interface{}, time int) error {
	conn := poll.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}
