package utils

import (
	"testing"
)

func Test_Base64(t *testing.T) {
	src := []byte("你好啊")
	result := Base64(src)
	t.Log(result)
}
