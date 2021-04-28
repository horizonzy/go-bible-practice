package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
		go fetch(url, ch)
	}
	s1 := <-ch
	s2 := <-ch

	if s1 != s2 {
		fmt.Println("s1 != s2")
	} else {
		fmt.Println("s1 == s2")
	}

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read response body error: %v", err)
		return
	}
	ch <- fmt.Sprintf("%s", bytes)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs %s\n", secs, url)
}
