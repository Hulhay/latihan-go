package main

import "fmt"

func main() {
	days := [7]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "ahad"}
	weekday := days[0:5]
	weekend := days[5:7]

	fmt.Println("weekend :", weekend)
	fmt.Println("weekday :")

	for i, day := range weekday {
		fmt.Printf("%d %s\n", i+1, day)
	}
}