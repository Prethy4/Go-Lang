package main

import "fmt"

func main() {
	var x = [5]int{2, 5, 7, 4, 9}
	x[4] = 8
	fmt.Println(x)

	var c = make([]int, 3, 10)
	c = append(c, 3)
	fmt.Println(c)

	var c1 = make([]string, 2, 10)
	c1 = append(c1, "dd")
	fmt.Println(c1)

	var c3 = map[string]string{"sajh": "Tap", "dagsj": "sagh", "asdasgh": "SAg"}
	fmt.Println(c3)

	var c2 = map[int]int{
		1: 3, 4: 2, 3: 1}
	fmt.Println(c2)

	var i int

	fmt.Scanln(&i)
	if i%3 == 0 && i%5 == 0 {
		fmt.Println("Divisible", i)
	} else {
		fmt.Println("Not divisible")
	}

}
