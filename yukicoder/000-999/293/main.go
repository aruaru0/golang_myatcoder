package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	A, B := getI(), getI()
	AA, BB := A, B
	a := make([]int, 0)
	for A > 0 {
		a = append(a, A%10)
		A /= 10
	}
	b := make([]int, 0)
	for B > 0 {
		b = append(b, B%10)
		B /= 10
	}
	if len(a) > len(b) {
		out(AA)
		return
	}
	if len(a) < len(b) {
		out(BB)
		return
	}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == 4 && b[i] == 7 {
			out(AA)
			return
		}
		if a[i] == 7 && b[i] == 4 {
			out(BB)
			return
		}
		if a[i] > b[i] {
			out(AA)
			return
		}
		if a[i] < b[i] {
			out(BB)
			return
		}
	}
	out(AA)
}
