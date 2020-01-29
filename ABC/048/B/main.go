package main

import (
	"fmt"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

func main() {

	var a, b, x uint64

	fmt.Scan(&a, &b, &x)

	var y uint64
	if a == 0 {
		y = b/x + 1
	} else {
		y = b/x - (a-1)/x
	}

	out(y)
}
