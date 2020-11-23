package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup(config *Config) error {
	var err error
	// db connect
	db, err = gorm.Open(config.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Name))

	if err != nil {
		return err
	}

	// set default table prefix
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)  // 设置空闲连接数
	db.DB().SetMaxOpenConns(100) // 设置最大打开连接数，默认值为0表示不限制

	return nil
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
