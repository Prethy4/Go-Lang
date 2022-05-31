package main

import "fmt"

func main() {
	var m1 = map[string]string{"un": "Prethy", "pwd": "12345678"}
	fmt.Println(m1)

	var m2 = map[string]int{}
	m2["un"] = 1
	m2["pwd"] = 1234
	fmt.Println(m2["un"], m2["pwd"])
}
