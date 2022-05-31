package main

import "fmt"

func main() {
	x := 20
	if x%2 == 0 && x%5 == 0 {
		fmt.Println("Divisible by 2 and 5")
	} else if x%2 == 0 && x%5 != 0 {
		fmt.Println("Not Divisible by 2 but not 5")
	} else {
		fmt.Println("Not Divisibl by ")
	}
}
