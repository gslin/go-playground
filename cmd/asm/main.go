package main

import "fmt"

func add(x, y int) int

func main() {
	a := 1
	b := 2
	c := add(a, b)

	fmt.Printf("%v + %v = %v\n", a, b, c)
}
