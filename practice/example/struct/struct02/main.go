package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	NameSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	AgeSet  = "0123456789"
)

func GenerateRandomStudentName(n int) string {
	name := make([]byte, n)
	for index := range name {
		name[index] = NameSet[rand.Intn(len(NameSet))]
	}
	return string(name)
}

func GenerateRandomAge() int {
	age := make([]byte, 2)
	for index := range age {
		age[index] = AgeSet[rand.Intn(len(AgeSet))]
	}
	ageVal, err := strconv.Atoi(string(age))
	if err != nil {
		panic("invalid age!! cannot convert into a number")
	}

	return ageVal
}

type Student struct {
	Name  string
	Score func(int) *int
	Age   int
}

func NewStudent(name string, op func(int) *int, age int) *Student {
	return &Student{
		Name:  name,
		Score: op,
		Age:   age,
	}
}

func (stu *Student) CalculateScore(num int) *int {
	return stu.Score(num)
}

func GetScore(num int) *int {
	return &num
}

func main() {
	stuA := NewStudent(GenerateRandomStudentName(5), GetScore, GenerateRandomAge())
	fmt.Printf("%+v\n", stuA)
	stuA.CalculateScore(5)
	fmt.Printf("%+v\n", stuA)
	// fmt.Println(studentA.Score(20))
}
