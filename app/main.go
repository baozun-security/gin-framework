package main

import (
	"baozun.com/framework/app/controllers"
	"baozun.com/framework/app/pkgs/logger"
	"baozun.com/framework/app/pkgs/mysql"
	"baozun.com/framework/app/pkgs/redis"
	"baozun.com/framework/app/pkgs/server"
	"baozun.com/framework/app/pkgs/setting"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"runtime"
)

var (
	runMode string // app run mode, available values are [dev|prd], default to dev
	cfgPath string // app config path
)

// 调整并发性能
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 解析命令参数
func initArgs() {
	// app --runMode prd
	flag.StringVar(&runMode, "runMode", "dev", "-runMode=[dev|prd]")
	flag.StringVar(&cfgPath, "cfgPath", "", "-cfgPath=/path/config/")
	flag.Parse()

	// verify run mode
	if mode := setting.RunMode(runMode); !mode.IsValid() {
		flag.PrintDefaults()
		return
	}

	// adjust config path
	if cfgPath == "" {
		var err error

		cfgPath, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	} else {
		cfgPath = path.Clean(cfgPath)
	}
}

// 初始化配置
func initConfig() {
	// setup config
	if err := setting.Setup(runMode, cfgPath); nil != err {
		log.Panicf("Faield to load config. %v\n", err)
	}
	// setup logger
	if err := logger.Setup(setting.Options.Logger); nil != err {
		log.Panicf("Faield to setup logger. %v\n", err)
	}
	// setup database
	if err := mysql.Setup(setting.Options.Database); nil != err {
		log.Panicf(fmt.Sprintf("Faild to setup database. %v\n", err))
	}
	// setup redis
	if err := redis.Setup(setting.Options.Redis); nil != err {
		log.Panicf(fmt.Sprintf("Faild to setup redis. %v\n", err))
	}
	// set gin mode
	gin.SetMode(setting.Options.Server.Mode)
}

func init() {
	initEnv()    // 初始化线程
	initArgs()   // 解析命令行参数
	initConfig() // 初始化配置
}

func main() {
	server.Run(controllers.InitRouter(), setting.Options.Server)
}
