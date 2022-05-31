package main

import "fmt"

func main() {
	var m3 = map[string]int{"a": 0, "b": 2}
	fmt.Println(m3)

	x, ok := m3["d"]
	fmt.Println(x, ok)
}
