package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")

	call("https://badssl.com/robots.txt")
	call("https://expired.badssl.com/robots.txt")
}
