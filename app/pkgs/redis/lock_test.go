package redis

import (
	"fmt"
	"testing"
)

func TestLock(t *testing.T) {
	lock, ok, err := TryLock("test_key", "test_value")
	if err != nil {
		fmt.Println("Error while attempting lock")
	}
	if !ok {
		fmt.Println("Lock failed")
	}
	fmt.Println("All keys: ", GetAllKeys())
	if nil != lock {
		fmt.Println("unLock result: ", lock.UnlockDeferDefault())
	}
}
