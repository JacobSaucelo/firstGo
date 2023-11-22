package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var link string = "https://dummyjson.com/posts"

func main() {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

// https://pkg.go.dev/net/http@go1.21.4 docccc
