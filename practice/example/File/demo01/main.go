package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open an existing file, then overwrite something else
	// O_TRUNC: clean up the file first
	output_file_path := "/home/xmzhao/practice/example/abc.txt"
	file, err := os.OpenFile(output_file_path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	// write contents
	str := "hello ximeng\r\n"
	//While writing, we use buffered writer. WriteString first writes into buffer then only do flush will write data to the disk.
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
