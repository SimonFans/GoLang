package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filePath := range os.Args[1:] {
		b, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		for _, line := range strings.Split(string(b), "\n") {
			counts[line] += 1
		}
	}
	for line, cnt := range counts {
		if cnt > 1 {
			fmt.Printf("%d\t%s\n", cnt, line)
		}
	}
}

// UT
// file1.txt
test1
test1
test12
test1
test13
test13
// file2.txt
test2
test21
test22
test23
test21
test2
// result
3       test1
2       test13
2       test2
2       test21
