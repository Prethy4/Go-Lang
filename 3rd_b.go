package main

import "fmt"

func main() {
	/*var x = make([]int, 0 ,4)
	x =append(x,1)
	x =append(x,2)
	x =append(x,3)*/
	var x = []int{1, 2, 3, 4}
	fmt.Println(x)

	var y = make([]int, 2)
	copy(y, x[1:3])
	fmt.Println(y)
}
