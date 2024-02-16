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

func combinations(list []int, k, buf int) (c chan []int) {
	c = make(chan []int, buf)
	n := len(list)

	pattern := make([]int, k)

	var body func(pos, begin int)
	body = func(pos, begin int) {
		if pos == k {
			t := make([]int, k)
			copy(t, pattern)
			c <- t
			return
		}

		for num := begin; num < n+pos-k+1; num++ {
			pattern[pos] = list[num]
			body(pos+1, num+1)
		}
	}
	go func() {
		defer close(c)
		body(0, 0)
	}()

	return
}

var N int
var a []int
var cnt int

// 与えられた3つの長さの線分で三角形が作れるかどうかを判定する関数
func isTriangle(a, b, c int) bool {
	// 三角形の成立条件：任意の2辺の長さの和が残りの1辺の長さよりも大きい
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

func hash(x []int) int {
	ret := 0
	for _, e := range x {
		ret = ret*37 + e
		ret %= 1e9 + 7
	}
	return ret
}

func rec(x []int) {
	if len(x) == 3 {
		i, j, k := x[0], x[1], x[2]
		if isTriangle(a[i], a[j], a[k]) {
			cnt++
		}
		return
	}
	for e := range combinations(x, 3, 1) {
		ii, jj, kk := e[0], e[1], e[2]
		if isTriangle(a[ii], a[jj], a[kk]) == false {
			continue
		}
		b := make([]int, 0)
		for i, j := 0, 0; j < len(x); j++ {
			if i < 3 && e[i] == x[j] {
				i++
			} else {
				b = append(b, x[j])
			}
		}
		// out(e, b)
		rec(b)
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	a = getInts(N * 3)

	x := make([]int, 0)
	for i := 0; i < 3*N; i++ {
		x = append(x, i)
	}
	rec(x)

	for i := N; i > 1; i-- {
		cnt /= i
	}
	out(cnt)
}
