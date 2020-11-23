package logger

import (
	"errors"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Output        string `yaml:"output"`
	Level         string `yaml:"level"`
	Format        string `yaml:"format"`
	Filename      string `yaml:"filename"`
	RotationCount uint   `yaml:"rotation_count"`

	// internal
	loggerLevel    logrus.Level
	loggerFormat   logrus.Formatter
	loggerFilename string
}

func (c *Config) parseFormat(format string) (logrus.Formatter, error) {
	if 0 == len(c.Format) {
		return &logrus.JSONFormatter{}, nil
	}
	switch strings.ToUpper(c.Format) {
	case "JSON":
		return &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcName, filename + ":" + strconv.Itoa(f.Line)
			},
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "@timestamp",
				logrus.FieldKeyFile: "@location",
				logrus.FieldKeyFunc: "@caller",
			},
		}, nil
	case "TEXT":
		return &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcName, filename + ":" + strconv.Itoa(f.Line)
			},
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "@timestamp",
				logrus.FieldKeyFile: "@location",
				logrus.FieldKeyFunc: "@caller",
			},
		}, nil
	}
	return nil, fmt.Errorf("not a valid logrus Format: %s", format)
}

func (c *Config) FillWithDefaults() error {
	var err error
	if c == nil {
		return errors.New("config is nil")
	}

	// adjust log level
	if c.loggerLevel, err = logrus.ParseLevel(c.Level); nil != err {
		return err
	}

	// adjust log format
	if c.loggerFormat, err = c.parseFormat(c.Format); nil != err {
		return err
	}

	// adjust output
	c.loggerFilename = c.Output
	if 0 == len(c.loggerFilename) {
		if root, err := os.Getwd(); err == nil {
			c.loggerFilename = root
		}
	}

	suffix := ".log"
	if !strings.HasSuffix(c.loggerFilename, suffix) {
		if 0 == len(c.Filename) {
			c.Filename = filepath.Base(strings.ToLower(c.Output))
		}

		if !strings.HasSuffix(c.Filename, suffix) {
			c.Filename += suffix
		}
		c.loggerFilename = filepath.Join(c.loggerFilename, c.Filename)
	}

	return nil
}

func (c *Config) NewWriter() (w io.Writer, err error) {
	switch strings.ToUpper(c.Output) {
	case "STDOUT":
		w = os.Stdout
	case "STDERR":
		w = os.Stderr
	case "NIL", "NULL":
		w = ioutil.Discard
	default:
		w, err = rotatelogs.New(
			c.loggerFilename+".%Y%m%d",
			//rotatelogs.WithLinkName(c.loggerFilename), // 生成软链，指向最新日志文件
			rotatelogs.WithRotationCount(c.RotationCount),            // 设置文件清理前最多保存的个数
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), // 设置日志分割的时间
		)
	}
	return
}
