package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	year int
}

func (p person) sayHello() {
	fmt.Println("Hello", p.name, "!")
}

func (p person) getAge() int {
	now := time.Now().Year()
	age := now - p.year
	return age
}

func main() {
	pras := person{"Prasetyo", 1998}
	pras.sayHello()

	prasAge := pras.getAge()
	fmt.Println("Now,", pras.name, "is", prasAge, "years old")
}