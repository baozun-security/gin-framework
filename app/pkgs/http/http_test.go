package http

import (
	"fmt"
	"testing"
)

func TestHttp(t *testing.T) {
	client := NewClient(nil)
	bytes, err := client.SetBaseUrl("http://10.101.191.106:5000/v1").Get("kri")
	if nil != err {
		fmt.Println("request err: ", err)
	} else {
		fmt.Println("request success. result: ", string(bytes))
	}
}
