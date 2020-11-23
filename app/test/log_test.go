package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_log(t *testing.T) {
	// 定义一个文件
	fileName := "ll.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	w := ioutil.Discard
	// 创建一个日志对象
	debugLog := log.New(w, "[Debug]", log.LstdFlags)
	debugLog.Println("A debug message here")
	//配置一个日志格式的前缀
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here ")
	//配置log的Flag参数
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")
}

func Test_02(t *testing.T) {
	path, _ := os.Getwd()
	fmt.Println("path >>>>>>>>>>>>>>", path)

	suffix := ".log"
	res := strings.HasSuffix(path + ".log", suffix)
	fmt.Println("res: ", !res)

	filename := filepath.Base(strings.ToLower(path))
	fmt.Println("filename >>>>>>>>>>>>>>>>>>>", filename)
}
