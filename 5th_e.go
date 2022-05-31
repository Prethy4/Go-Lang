package main

import "fmt"

func main() {
	var flag bool = true
	var x = 5
	for flag {
		x++
		if x == 10 {
			flag = false
		}
		fmt.Println("I want work in Google.")
	}
}
