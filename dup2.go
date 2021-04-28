package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:]

	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open file error, %v\n", err)
			continue
		}
		checkFileIsUniq(file)
	}

}

func checkFileIsUniq(file *os.File) {

	var before string
	input := bufio.NewScanner(file)
	for input.Scan() {
		currentLine := input.Text()
		if before == currentLine {
			fmt.Printf(file.Name())
			break
		}
		before = currentLine
	}
}
