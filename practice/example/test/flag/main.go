package main

import (
	"flag"
	"fmt"
)

func main() {
	firstNamePtr := flag.String("firstname", "Ximeng", "give a first name")
	var lastName string
	flag.StringVar(&lastName, "lastname", "Zhao", "give a last name")
	flag.Parse()
	fmt.Println("first name: ", *firstNamePtr)
	fmt.Println("last name: ", lastName)
}
