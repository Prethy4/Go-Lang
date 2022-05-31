package main

import "fmt"

func main() {
	var x = make([]int, 1, 2)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 2)
	fmt.Println(x, len(x), cap(x))
}
