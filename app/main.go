package main

import (
	"baozun.com/leak/app/controllers"
	"baozun.com/leak/app/pkgs/logger"
	"baozun.com/leak/app/pkgs/mysql"
	"baozun.com/leak/app/pkgs/redis"
	"baozun.com/leak/app/pkgs/setting"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"time"
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
}

func main() {
	// 初始化线程
	initEnv()
	// 解析命令行参数
	initArgs()
	// 初始化配置
	initConfig()

	gin.SetMode(setting.Options.Server.Mode)
	endPoint := fmt.Sprintf("%s:%d", setting.Options.Server.Addr, setting.Options.Server.Port)
	server := &http.Server{
		Addr:         endPoint,
		Handler:      controllers.Init(), // 初始化控制器,
		ReadTimeout:  setting.Options.Server.ReadTimeout * time.Second,
		WriteTimeout: setting.Options.Server.WriteTimeout * time.Second,
	}
	logger.Logger.Infof("start http server listening %s", endPoint)

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
