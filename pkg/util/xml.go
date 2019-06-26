package util


import (
	
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
)

func ReadSource(src, pwd string) ([]byte, error) {

	abspath := filepath.Join(src, pwd)
	file, err := os.Open(abspath)
	defer file.Close()
	
	if err != nil {
		fmt.Println("文件读取错误")
	}

	fmt.Println(abspath)
	return ioutil.ReadAll(file)
}