package logger

import (
	"github.com/sirupsen/logrus"
	"io"
)

var Logger = logrus.New()

// Setup initialize the log instance
func Setup(config *Config) error {
	var (
		err    error
		writer io.Writer
	)
	err = config.FillWithDefaults()
	if nil != err {
		return err
	}

	writer, err = config.NewWriter()
	if nil != err {
		return err
	}
	Logger.SetOutput(writer)
	Logger.SetLevel(config.loggerLevel)
	Logger.SetFormatter(config.loggerFormat)
	//Logger.AddHook(&AppHook{appName})
	Logger.SetReportCaller(true)

	return nil
}
