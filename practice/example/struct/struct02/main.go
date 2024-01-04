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

type Student struct {
	Name  string
	Score func(int) *int
	Age   int
}

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

func GetScore(x int) func(int) *int {
	return func(x int) *int {
		result := x * 2
		return &result
	}
}

func main() {
	studentA := Student{
		Name:  GenerateRandomStudentName(5),
		Age:   GenerateRandomAge(),
		Score: GetScore(10),
	}
	fmt.Println(studentA.Score(20))
}
