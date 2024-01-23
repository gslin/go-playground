package main

import (
	"fmt"

	"github.com/gslin/go-playground/internal/https"
)

func main() {
	fmt.Println("Hello, world.")

	https.Get("https://badssl.com/robots.txt")
	https.Get("https://expired.badssl.com/robots.txt")
}
