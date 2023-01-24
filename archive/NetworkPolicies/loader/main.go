package main

import (
	"fmt"
	"net/http"
)

func main() {
	for {

		requestURL := "http://20.126.224.206:9091"
		res, err := http.Get(requestURL)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)

	}
}
