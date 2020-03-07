package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

const eps = 10e-9

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()

LOOP:
	for h := 1; h <= 3500; h++ {
		for n := 1; n <= 3500; n++ {
			a := h * N * n
			b := 4*h*n - h*N - n*N
			if b != 0 {
				w := a / b
				if a%b == 0 && w > 0 {
					out(h, n, w)
					break LOOP
				}
			}
		}
	}
}
