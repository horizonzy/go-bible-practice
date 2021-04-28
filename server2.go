package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler1(write http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println("enter handler")
	fmt.Fprintf(write, "URL.Path = %q\n", request.URL.Path)
}

func counter(write http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Println("enter counter")
	fmt.Fprintf(write, "Count %d\n", count)
	mu.Unlock()
}
