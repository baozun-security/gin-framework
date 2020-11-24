package server

import (
	"fmt"
	"time"
)

// server config
type Config struct {
	Addr         string        `yaml:"addr"`
	Port         int           `yaml:"port"`
	Mode         string        `yaml:"mode"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

// Name returns client name of the config
func (c *Config) Endpoint() string {
	return fmt.Sprintf("%s:%d", c.Addr, c.Port)
}
