package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	wordCounts := make(map[string]int)
	for _, filePath := range os.Args[1:] {
		b, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		for _, line := range strings.Split(string(b), "\n") {
			for _, word := range strings.Split(line, " ") {
				wordCounts[word] += 1
			}
		}
	}
	fmt.Println(wordCounts)
}

// % go run main.go /Users/.../go/src/helloworld/files/test1.txt
// test1.txt
Java Python
Python
Python
Golang
Golang
SQL Golang Python
// map[Golang:3 Java:1 Python:4 SQL:1]
