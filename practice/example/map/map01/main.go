package main

import "fmt"

func main() {
	monsters := make([]map[string]string, 2)
	fmt.Println(monsters)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "Mochi"
		monsters[0]["age"] = "3"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "Kopi"
		monsters[1]["age"] = "2"
	}
	fmt.Println(monsters)
	// use slice's append function to add new map
	newMonster := map[string]string{
		"name": "apple",
		"age":  "5",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
