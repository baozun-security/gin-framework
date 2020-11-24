package setting

import (
	"baozun.com/leak/app/pkgs/logger"
	"baozun.com/leak/app/pkgs/mysql"
	"baozun.com/leak/app/pkgs/redis"
	"baozun.com/leak/app/pkgs/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Find config file based on run mode.
var FindModeConfigFile = func(mode, cfgPath string) string {
	cfgPath = filepath.Clean(cfgPath)

	// adjust cfgPath
	switch cfgPath {
	case ".":
		cfgPath, _ = os.Getwd()
	case "..":
		var err error

		cfgPath, err = os.Getwd()
		if err == nil {
			cfgPath = filepath.Dir(cfgPath)
		}
	}

	return filepath.Join(cfgPath, "config", RunMode(mode).String()+".yml")
}

// app config // 自定义的一些配置
type App struct {
	Name string `yaml:"name"`
}

// AppOptions defines specs for config
type AppOptions struct {
	App      *App           `yaml:"app"`
	Logger   *logger.Config `yaml:"logger"`
	Server   *server.Config `yaml:"server"`
	Database *mysql.Config  `yaml:"database"`
	Redis    *redis.Config  `yaml:"redis"`
}

// Shared Options
var Options *AppOptions

// Setup initialize the configuration instance
func Setup(runMode, cfgPath string) error {
	filename := FindModeConfigFile(runMode, cfgPath)

	// parse yaml config
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &Options)
	if err != nil {
		return err
	}
	return nil
}
