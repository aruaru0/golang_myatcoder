package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	sc.Split(bufio.ScanWords)

	N, K := float64(getInt()), float64(getInt())

	a := math.Pow(1/N, 3)
	b := (K - 1) * (N - K) * 6
	c := (K - 1) * 3
	d := (N - K) * 3

	out((b + c + d + 1) * a)
}
