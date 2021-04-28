package main

import "fmt"

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}

func main() {
	fmt.Println(Signum(0))
}

type Point struct {
	x, y int
	z    string
}
