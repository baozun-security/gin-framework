package utils

import "reflect"

// 获取数据类型
func GetType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
