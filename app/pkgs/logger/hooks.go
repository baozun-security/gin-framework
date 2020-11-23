package logger

import (
	"github.com/sirupsen/logrus"
)

type AppHook struct {
	appName string
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = h.appName
	return nil
}
