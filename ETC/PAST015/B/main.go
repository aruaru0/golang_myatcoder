package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var h, c int
	fmt.Scan(&c, &h)

	if max(c, h) >= 2800 {
		fmt.Println("o")
	} else {
		fmt.Println("x")
	}
}
