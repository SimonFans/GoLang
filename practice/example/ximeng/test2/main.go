package main

import "fmt"

type People struct {
	Name  string
	Age   int
	Skill *SuperStar
}

type SuperStar struct {
	SkillName string
	Power     int
}

func main() {
	fmt.Println("----code starts here ----")
	super1 := SuperStar{SkillName: "Dugujiujian", Power: 500}
	fmt.Println(super1)
	p1 := People{Name: "Xiaoming", Age: 30, Skill: &super1}
	fmt.Println(p1.Name, p1.Age, p1.Skill.SkillName, p1.Skill.Power)
}
