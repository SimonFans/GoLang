package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readPath := "/home/xmzhao/mygit/GoLang/practice/example/abc.txt"
	writepath := "/home/xmzhao/mygit/GoLang/practice/example/efg.txt"
	// Read file content
	data, err := ioutil.ReadFile(readPath)
	if err != nil {
		fmt.Printf("read file error=%v\n", err)
		return
	}
	// overwrite operation
	err = ioutil.WriteFile(writepath, data, 0666)
	if err != nil {
		fmt.Printf("read file error=%v\n", err)
	}

}
