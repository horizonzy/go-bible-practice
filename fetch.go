package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch error: %v", err)
			return
		}
		_, xx := io.Copy(os.Stdout, resp.Body)
		if xx != nil {
			fmt.Fprintf(os.Stderr, "copy response error: %v", err)
			return
		}
	}
}
