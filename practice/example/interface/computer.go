package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Usb interface {
	// declare two mei you shi xian de method
	Start()
	Stop()
}

// shi xian le interface all methods then you can do var a Speak = <zi ding yi bian liang name>
type Speak interface {
	SayHi()
}

type Student1 struct {
	Name string
}

func (stu Student1) SayHi() {
	fmt.Printf("Hello %v\n", stu.Name)
}

type Phone struct {
}

func (p *Phone) Start() {
	fmt.Println("Phone is using the USB interface")
}

func (p *Phone) Stop() {
	fmt.Println("Phone stops using the USB interface")
}

type Camera struct {
}

func (p *Camera) Start() {
	fmt.Println("Camera is using the USB interface")
}

func (p *Camera) Stop() {
	fmt.Println("Camera stops using the USB interface")
}

type Computer struct {
}

func (c *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

// hero implements func Sort(data Interface)
type Student struct {
	Name  string
	Age   int
	Score float64
}

// declare a slice with Hero struct type
type StuSlice []Student

// Implement Interface jie kou cai neng yong Sort() method
func (s StuSlice) Len() int {
	return len(s)
}

func (s StuSlice) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s StuSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(&phone)
	computer.Working(&camera)
	stu := Student1{"Ximeng"}
	var si Speak = stu
	si.SayHi()
	fmt.Println("***********")
	var students StuSlice
	for i := 0; i < 10; i++ {
		student := Student{
			Name:  fmt.Sprintf("Student~%d", rand.Intn(100)),
			Age:   rand.Intn(30),
			Score: rand.Float64(),
		}
		students = append(students, student)
	}
	// Before sorting
	for _, v := range students {
		fmt.Println(v)
	}
	// sort by score
	sort.Sort(students)
	fmt.Println("After sorting.......")
	// After sorting
	for _, v := range students {
		fmt.Println(v)
	}
}
