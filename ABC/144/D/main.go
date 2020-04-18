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

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	sc.Split(bufio.ScanWords)

	a, b, x := float64(getInt()), float64(getInt()), float64(getInt())

	h := 2 * x / (a * b)
	s := 2 * (a*a*b - x) / (a * a)

	ans := math.Atan2(b, h) * 180 / math.Pi
	if h > a {
		ans = math.Atan2(s, a) * 180 / math.Pi
	}

	// out(b, h)
	// out(a, s)
	fmt.Printf("%-12.12f\n", ans)
}
