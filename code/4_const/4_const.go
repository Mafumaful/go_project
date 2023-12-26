package main

import (
	"fmt"
)


// iota can only be used with constants
func main() {
	const N = 5

	//print the value of N
	fmt.Println(N)

	const(
		// iota is a special constant, it starts with 0 and increments by one
		APPLE = iota
		ORANGE
		BANANA
	)

	fmt.Println("APPLE = ", APPLE)
	fmt.Println("ORANGE = ", ORANGE)
	fmt.Println("BANANA = ", BANANA)

	const(
		// iota can be used to skip values
		RED = iota
		BLUE
		_
		YELLOW
	)

	fmt.Println("RED = ", RED)
	fmt.Println("BLUE = ", BLUE)
	fmt.Println("YELLOW = ", YELLOW)

	const(
		a, b = iota + 1, iota + 2
		c, d
		e, f

		g, h = iota * 2, iota * 3
	)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	
}
