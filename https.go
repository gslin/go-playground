package main

import (
	"fmt"
	"net/http"
)

func call(url string) {
	res, err := http.Get(url)
	fmt.Printf("url: %s\n", url)
	fmt.Printf("res: %s\n", res)
	fmt.Printf("err: %s\n", err)
}
