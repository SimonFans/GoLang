package main

import "fmt"

// Testing center
type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("Student Name=%v, Age=%v, Score=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

// Pupil
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("Pupil student is testing...")
}

// Grdduate
type Graduate struct {
	Student
}

func (p *Graduate) testing() {
	fmt.Println("Graduate student is testing...")
}

func main() {
	pupil := &Pupil{}
	pupil.Name = "Mochi"
	pupil.Age = 1
	pupil.testing()
	pupil.ShowInfo()
	pupil.SetScore(50)
	pupil.ShowInfo()
	graduate := &Graduate{}
	graduate.Name = "Kopi"
	graduate.Age = 10
	graduate.testing()
	graduate.SetScore(20)
	graduate.ShowInfo()
}
