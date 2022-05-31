package main

import "fmt"

func main() {
	var x = []int{5, 10, 15}
	for _, v := range x {
		fmt.Println(v)
	}
}
