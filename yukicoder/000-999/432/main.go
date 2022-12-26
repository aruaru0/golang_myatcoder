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

/*
	X := []string{"a", "b", "c", "d", "e", "f", "g", "i"}
	x := make([]string, len(X))
	for i := 0; i < len(X); i++ {
		x[i] = X[i]
	}
	for len(x) != 1 {
		y := make([]string, 0)
		for i := 1; i < len(x); i++ {
			y = append(y, x[i-1]+x[i])
		}
		x = y
	}
	out(x, len(X))
	for _, e := range X {
		out(string(e), strings.Count(x[0], string(e)))
	}
*/

func f() {
	s := getS()
	N := len(s)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = int(s[i] - '0')
	}

	for len(a) != 1 {
		b := make([]int, len(a)-1)
		for i := 0; i < len(a)-1; i++ {
			x := a[i] + a[i+1]
			b[i] = (x/10 + x%10) % 10
		}
		// out(b)
		a = b
	}
	out(a[0])
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		f()
	}
}
