package main

import (
	"fmt"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

func main() {

	var x uint64
	fmt.Scan(&x)

	m := x / 11
	r := x % 11
	m *= 2
	if r == 0 {

	} else if r <= 6 {
		m++
	} else {
		m += 2
	}
	out(m)

}
