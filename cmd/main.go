package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello vim-go")
	sum := add(5, 10)
	fmt.Println(sum)
}

func add(x, y int) int {
	return x + y
}
