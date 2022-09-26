package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	// If path exists then err = nil
	if err == nil {
		return true, nil
	}
	// If returned error type was true by calling os.IsNotExist, then file or folder not exist
	if os.IsNotExist(err) {
		return false, nil
	}
	//  If returned error is other type then uncertain if the file path exists or not
	return false, err
}

func main() {
	path := "/home/xmzhao/mygit/GoLang/practice/example/abc.txt"
	fmt.Println(PathExists(path))
}
