package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Skill string
}

func NewPerson() {
	person := Person{
		Name:  "Wang lu",
		Age:   20,
		Skill: "abc",
	}
	// struct -> json serialize
	data, err := json.Marshal(&person)
	if err != nil {
		fmt.Printf("Serialize error=%v", err)
	}
	// output after serialization
	fmt.Printf("After serialization=%v, the type is=%T\n", string(data), string(data))
}

func NewMap() {
	aMap := make(map[string]interface{})
	aMap["name"] = "xiao wang"
	aMap["address"] = "123th st"
	aMap["zipcode"] = 98052
	// map -> json serialize
	data, err := json.Marshal(aMap)
	if err != nil {
		fmt.Printf("Serialize error=%v", err)
	}
	// output after serialization
	fmt.Printf("After serialization=%v, the type is=%T\n", string(data), string(data))
}

// define a slice with multiple key/value pairs
func NewSlice() {
	var slice []map[string]interface{}
	// Need to run make before using map
	m1 := make(map[string]interface{})
	m1["name"] = "xiao li"
	m1["address"] = []string{"US", "CH"}
	m1["zipcode"] = 92050
	m2 := make(map[string]interface{})
	m2["name"] = "Tom"
	m2["address"] = "3th st"
	m2["zipcode"] = 96000
	slice = append(slice, m1)
	slice = append(slice, m2)
	// slice -> json serialization
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("Serialize error=%v", err)
	}
	// output after serialization
	fmt.Printf("After serialization=%v, the type is=%T\n", string(data), string(data))
}

func main() {
	// serilize (struct) -> json string
	NewPerson()
	// serilize (map) -> json string
	NewMap()
	// serialize (slice) -> json string
	NewSlice()
}
