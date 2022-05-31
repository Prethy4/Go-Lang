package main

import "fmt"

func main() {
	//var x = [...]int{1, 2, 3}
	//x = [3] int {4,5,6}
	var x = []int{1, 2, 3, 4}
	fmt.Println(x)

	var y = make([]int, 2)
	copy(y, x[2:4])
	fmt.Println(y)
}
