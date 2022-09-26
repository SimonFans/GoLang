package main

import (
	"flag"
	"fmt"
)

func main() {
	// define variables to get command line params
	// Use flag package, when you enter the input, no need to follow the order, very flexible.
	var (
		user string
		pwd  string
		host string
		port int
	)
	flag.StringVar(&user, "u", "", "Username")
	flag.StringVar(&pwd, "pwd", "", "Password")
	flag.StringVar(&host, "h", "", "Host Name")
	flag.IntVar(&port, "port", 3396, "Port Number")
	flag.Parse()
	// output
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", user, pwd, host, port)
}
