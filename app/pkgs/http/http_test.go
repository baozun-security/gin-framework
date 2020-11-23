package http

import (
	"fmt"
	"testing"
)

func TestHttp(t *testing.T) {
	client := NewClient(nil)
	bytes, err := client.Get("http://10.101.191.106:5000/v1/kri")
	if nil != err {
		fmt.Println("request err: ", err)
	} else {
		fmt.Println("request success. result: ", string(bytes))
	}
}
