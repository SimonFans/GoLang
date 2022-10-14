package main

import (
	"fmt"
	"reflect"
)

type T struct {
	Name string "yaml:\"1\""
}

type S struct {
	Color string `color:"Blue" model:"Honda"`
}

func main() {
	// t is a Type
	t := reflect.TypeOf(T{})
	fmt.Printf("Type=%v, Kind=%v\n", t, t.Kind())
	// Type has a method FieldByName returns the struct field with the given name
	// FieldByName return a StructField struct which includes many fields you can call such as Tag
	f, _ := t.FieldByName("Name")
	nameTagVal := f.Tag.Get("yaml")
	fmt.Println(nameTagVal)
	fmt.Println(reflect.TypeOf(nameTagVal))

	// S is a Type
	s := reflect.TypeOf(S{})
	fmt.Printf("Type=%v, Kind=%v\n", s, s.Kind())
	// return structfield -> call .Tag => color:"Blue" model:"Honda"
	f1 := s.Field(0).Tag
	fmt.Println(f1)
	fmt.Printf("color=%v, model=%v\n", f1.Get("color"), f1.Get("model"))
}
