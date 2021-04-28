package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler2)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler2(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(write, "%s %s %s\n", request.Method, request.URL, request.Proto)
	for k, v := range request.Header {
		fmt.Fprintf(write, "Header [%q] = %q\n", k, v)
	}
	fmt.Fprintf(write, "Host = %q\n", request.Host)
	fmt.Fprintf(write, "RemoteAddr = %q\n", request.RemoteAddr)
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range request.Form {
		fmt.Fprintf(write, "From [%q] = %q\n", k, v)
	}
}
