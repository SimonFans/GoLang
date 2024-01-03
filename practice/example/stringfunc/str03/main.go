package main

import "fmt"

func removeString(slice []string, s string) []string {
	var result []string
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return result
}

func main() {
	newSlice := removeString([]string{"abc", "aad", "ddd"}, "aad")
	fmt.Println(newSlice)
}
