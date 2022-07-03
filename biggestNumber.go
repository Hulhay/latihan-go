package main

import "fmt"

func biggestNumber(num1, num2, num3 int) int {
	if num1 > num2 && num1 > num3 {
		return num1
	} else if num2 > num1 && num2 > num3 {
		return num2
	} else {
		return num3
	}
}

func main() {
	fmt.Println(biggestNumber(10, 1, 24))
}