package http

import "time"

const (
	ConnectTimeout = 3 * time.Second
	RequestTimeout = 100 * time.Second
	KeepAliveTime  = 30 * time.Second
	AcceptContent  = "application/json;charset=utf-8"
	AcceptEncoding = "gzip"
)
