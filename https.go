package main

import (
	"fmt"
	"net/http"
)

func call(url string) {
	res, err := http.Get(url)
	fmt.Printf("%s: %s\n", res, err)
}
