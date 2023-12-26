package main

import (
	"fmt"
)

func fool(a string, b int) (string, int) {
	return a, b
}

func fool2(a string, b int) (k1 string, k2 int) {

	// k1 and k2 are already declared
	fmt.Println("k1 = ", k1)
	fmt.Println("k2 = ", k2)

	k1 = a
	k2 = b
	return
}

func main() {
	a, b := fool("foo", 5)
	fmt.Println(a, b)

	c, d := fool2("bar", 10)
	fmt.Println(c, d)
}