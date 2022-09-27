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

func NewPerson() string {
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
	return string(data)
}

func NewMap() string {
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
	return string(data)
}

// define a slice with multiple key/value pairs
func NewSlice() string {
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
	return string(data)
}

func unmarshalStruct() {
	str := NewPerson()
	var person Person
	// deserialize (JSON string to Struct type)
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Println("Deserialze to Struct => person: ", person)
}

func unmarshalMap() {
	str := NewMap()
	var m map[string]interface{}
	// deserialize (JSON string to Struct type)
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Println("Deserialze to Struct => Map: ", m)
}

func unmarshalSlice() {
	str := NewSlice()
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Println("Deserialze to Struct => Slice: ", slice)
}

func main() {
	// serilize (struct) -> json string
	NewPerson()
	// serilize (map) -> json string
	NewMap()
	// serialize (slice) -> json string
	NewSlice()
	// deserialzie (JSON string) -> Struct
	unmarshalStruct()
	// deserialzie (JSON string) -> Map
	unmarshalMap()
	// deserialzie (JSON string) -> Slice
	unmarshalSlice()
}
