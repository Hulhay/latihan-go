package main

import "fmt"

func main() {
	
	userData := []map[string]interface{}{
		{"name": "Prasetyo", "age": 24},
		{"name": "Dimas", "age": 23},
	}

	for _, user := range userData {
		fmt.Println(user)
	}

	// fmt.Println(user)
}