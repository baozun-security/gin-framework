package utils

import "time"

// 获取当前时间字符串
func GetInt64TimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 时间戳转换为时间字符串
func Int64TimeToString(value int64) string {
	return time.Unix(value, 0).Format("2006-01-02 15:04:05")
}
