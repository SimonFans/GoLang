package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)
	// set # CPU you want to use
	// runtime.GOMAXPROCS(cpuNum - 1)
}
