package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
OS.File (OS => package, File is a struct type)
*/

func main() {
	path := "/home/xmzhao/practice/example/items.csv"
	// open the file (file => address)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Open file error=", err)
	}
	fmt.Printf("file=%v\n", file)
	// close the file
	// if err = file.Close(); err != nil {
	// 	fmt.Println("Close file error=", err)
	// }
	defer file.Close()
	// <1> create a *BufferedReader. Default buffer size = 4096, fit for large file size
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

	// <2> create a one time reader. No need to open & close file, because it's encapsulated in the ReadFile function. Fit for small file size.
	file2 := "/home/xmzhao/practice/example/items.csv"
	content, err := ioutil.ReadFile(file2)
	if err != nil {
		fmt.Printf("Read file error=%v", err)
	}
	fmt.Printf("%v", string(content))
	lines := strings.Split(string(content), "\n")
	fmt.Println("lines=>", lines[1])

	// Write operation using BufferedWriter
	output_file_path := "/home/xmzhao/practice/example/abc.txt"
	file, err = os.OpenFile(output_file_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	// write contents
	str := "hello buddy\r\n"
	//While writing, we use buffered writer. WriteString first writes into buffer then only do flush will write data to the disk.
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
