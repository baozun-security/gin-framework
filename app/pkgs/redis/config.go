package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

// redis config
type Config struct {
	Host      string `yaml:"host"`
	Password  string `yaml:"password"`
	DB        int    `yaml:"database"`
	MaxIdle   int    `yaml:"max_idle"`
	MaxActive int    `yaml:"max_active"`
	Timeout   int    `yaml:"timeout"`
}

// Name returns client name of the config
func (c *Config) Name() string {
	return fmt.Sprintf("%s/%d", c.Host, c.DB)
}

// FillWithDefaults apply default values for fields with invalid values.
func (c *Config) FillWithDefaults() {
	if c.MaxIdle <= 0 {
		c.MaxIdle = DefaultPollMaxIdle
	}
	if c.MaxActive <= 0 {
		c.MaxActive = DefaultPollMaxActive
	}
	if c.Timeout <= 0 {
		c.Timeout = DefaultTimeout
	}
}

func (c *Config) Connect() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", c.Host,
		redis.DialPassword(c.Password),
		redis.DialDatabase(c.DB),
		redis.DialConnectTimeout(time.Duration(c.Timeout)*time.Second),
		redis.DialReadTimeout(time.Duration(c.Timeout)*time.Second),
		redis.DialWriteTimeout(time.Duration(c.Timeout)*time.Second))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
