package main

import "fmt"

func main() {
	var x int = 94
	var y int = 10
	fmt.Println("Data: ", x, y)

	//task1
	a := x + y
	fmt.Println("sum: ", a)

	b := x - y
	fmt.Println("sub:", b)

	c := x * y
	fmt.Println("multi", c)

	d := x / y
	fmt.Println("division: ", d)

	mod := x % y
	fmt.Println("Modulus:", mod)
	// taskend

	fmt.Println("For == :", x == y)
	fmt.Println("for !=", x != y)
	fmt.Println("for increment <<", x<<1)
	fmt.Println("for increment <<", y<<1)
	fmt.Println("for decrement >>:", x>>y)
	fmt.Println("for decrement >>", x>>y)

	// task2

	fmt.Println("For divisional of 3", x%3 == 0)
	fmt.Println("For divisional of 3 & 5", x%3 == 0 && x%5 == 0)
	fmt.Println("For divisional of 3 & 5", x != 0 && x%5 == 0)

	// endask

	if true {
		fmt.Println("hello!")
	}

	var f float64 = 10.5
	var i int = 10
	var z float64 = f + float64(i)
	fmt.Println(z)

	var flo float64 = 10.5
	var r float64 = 10.5
	fmt.Println(flo == r)

	var flo2 float64 = 10.5
	var r2 float64 = 105 / float64(10)
	fmt.Println(flo2 == r2)

	var flo3 float64 = 10.5
	var r3 float64 = 105 / 10
	fmt.Println(flo3 == r3)
}
