package util

import (
	"fmt"
	"testing"
)

func Test_SHA256(t *testing.T) {
	fmt.Println(SHA256(Md5("admin") + "8779870"))
}
