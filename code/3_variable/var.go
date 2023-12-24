package main

import "fmt"

func main() {
	var a int = 10
	var st string = "the number of a is:"
	fmt.Println(st, a)
	fmt.Printf("%s %d\n", st, a)

	b := 20
	fmt.Println(st, b)
	fmt.Printf("The format of d is %T\n", b)

	var c, d int = 30, 40
	fmt.Println(st, c, d)

	var (
		e int = 50
		f int = 60
	)
	fmt.Println(st, e, f)
}
