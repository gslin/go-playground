package https

import (
	"fmt"
	"net/http"
)

func Get(url string) {
	res, err := http.Get(url)
	fmt.Printf("url: %v\n", url)
	fmt.Printf("res: %v\n", res)
	fmt.Printf("err: %v\n", err)
}
