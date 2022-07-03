package main

import "fmt"

func split(num int) (ans1, ans2 int) {
	ans1 = num * 4 / 9
	ans2 = num - ans1
	return ans1, ans2
}

func main() {
	fmt.Println(split(17))
}