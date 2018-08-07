package main

import (
	"fmt"
)

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	fmt.Println("Hello World from GO")
	w1, w2 := "Hello", "Welcome to Go programming"
	fmt.Println(multiple(w1, w2))

	x := 40
	a := &x

	fmt.Println(a)
	fmt.Println(*a)

	*a = 60
	fmt.Println(x)
}
