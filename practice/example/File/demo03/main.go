package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	   task1: read and show in the terminal.
	   task2: append 5 strings to the existing file
	   os.O_RDWR => read/write mode open
	*/
	output_file_path := "/home/xmzhao/practice/example/abc.txt"
	file, err := os.OpenFile(output_file_path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	// Read current file
	reader := bufio.NewReader(file)
	for {
		// read row by row. du dao yi ge huan hang jiu jie shu
		str, err := reader.ReadString('\n')
		// read till the end of line
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	// write contents
	str := "Hello Beijing\r\n"
	//While writing, we use buffered writer. WriteString first writes into buffer then only do flush will write data to the disk.
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
