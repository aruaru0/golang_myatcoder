package main

import (
	"bufio"
	"fmt"
	"math"
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

const inf = int(1e18)

func f(a, b []int) int {
	if len(a) == 0 || len(b) == 0 {
		return inf
	}
	// out(a, b)
	ret := inf
	for i := 0; i < len(a); i++ {
		l := lowerBound(b, a[i])
		// out(l, a[i], "b", b[max(0, l-1)], b[min(len(b)-1, l)])
		ret = nmin(ret, abs(a[i]-b[min(len(b)-1, l)]),
			abs(a[i]-b[max(0, l-1)]))
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	r, g, b := 0, 0, 0
	R := make([]int, 0)
	G := make([]int, 0)
	B := make([]int, 0)
	for i := 0; i < 2*N; i++ {
		a, c := getI(), []byte(getS())[0]
		if c == 'R' {
			r++
			R = append(R, a)
		} else if c == 'G' {
			g++
			G = append(G, a)
		} else {
			b++
			B = append(B, a)
		}
	}
	r %= 2
	g %= 2
	b %= 2

	sort.Ints(R)
	sort.Ints(G)
	sort.Ints(B)

	if r == 0 && g == 0 && b == 0 {
		out(0)
		return
	}
	if r == 0 {
		out(min(f(G, B), f(R, G)+f(R, B)))
		return
	}
	if g == 0 {
		out(min(f(R, B), f(R, G)+f(G, B)))
		return
	}
	if b == 0 {
		out(min(f(R, G), f(R, B)+f(G, B)))
		return
	}
}
