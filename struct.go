package main

import "fmt"

type user struct {
	name string
	age int
}

func main() {
	users := []user{
		{name:"Prasetyo", age:24},
		{name:"Dimas", age:23},
	}

	for _, userdata := range users {
		fmt.Println(userdata.name, "is", userdata.age, "years old")
	}
}