package main

import "fmt"

func main() {
	var m1 = map[string]string{"un": "Prethy", "pwd": "12345678"}
	fmt.Println(m1)

	var m2 = map[int]int{
		0: 0, 1: 10, 2: 20,
	}
	fmt.Println(m2[0], m2[1], m2[2])
}
