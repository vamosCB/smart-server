package utils

import (
	"io/ioutil"
	"os"
)

// GetFileContent 获取文件内容
func GetFileContent(file string) ([]byte, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	content, err := ioutil.ReadAll(fd)
	return content, nil
}
