package main

import "fmt"

func main() {
	type student struct {
		firstName  string
		familyName string
		id         int
	}
	//x := student{"Tanjum", "Amin Prethy", 1}
	x := student{familyName: "Tanjum", firstName: "Amin Prethy", id: 1}
	fmt.Println(x)
}
