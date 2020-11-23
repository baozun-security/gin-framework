package test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func Test_logrus(t *testing.T) {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel) // "DEBUG", "INFO", "WARN", "ERROR", "FATAL"
	logger = log.WithFields(log.Fields{"request_id": "123444", "user_ip": "127.0.0.1"})
	logger.Info("hello, logrus....")
	logger.Info("hello, logrus1....")
	logger.Debug("this is a test .... ")
	// log.WithFields(log.Fields{
	// "animal": "walrus",
	// "size":  10,
	// }).Info("A group of walrus emerges from the ocean")

	// log.WithFields(log.Fields{
	// "omg":  true,
	// "number": 122,
	// }).Warn("The group's number increased tremendously!")

	// log.WithFields(log.Fields{
	// "omg":  true,
	// "number": 100,
	// }).Fatal("The ice breaks!")
}
