package main

import "fmt"

func main() {
	var x = []int{1, 2, 3}
	x = append(x, 4)
	x = append(x, 5)
	fmt.Println(x)

	//var y [][]int
	//y[0] = []int{1, 2, 3}
	//fmt.Println(y)
}
