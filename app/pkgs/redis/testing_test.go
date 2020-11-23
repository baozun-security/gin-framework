package redis

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

var (
	fakeConfig = &Config{
		Host:     "10.101.191.106:6379",
		Password: "",
		DB:       0,
	}
)

func TestMain(m *testing.M) {
	// setup redis
	if err := Setup(fakeConfig); nil != err {
		log.Panicf(fmt.Sprintf("Faild to setup redis. %v\n", err))
	}
	code := m.Run()
	os.Exit(code)
}
