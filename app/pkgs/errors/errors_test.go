package errors

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	fmt.Println("err: ", InvalidRequest().AttachMessage("Has a error.").Error())
}
