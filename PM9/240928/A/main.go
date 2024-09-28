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

func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	const MOD = 1000000007

	// var h, w int
	// fmt.Scan(&h, &w)
	h, w := getI(), getI()
	if h > w {
		h, w = w, h
	}
	a := make([]int, 0)
	for i := 0; i < 1<<h; i++ {
		ok := true
		for j := 1; j <= i; j <<= 1 {
			if (i&j) != 0 && (i&(j<<1)) != 0 {
				ok = false
				break
			}
		}
		if ok {
			a = append(a, i)
		}
	}
	n := len(a)
	b := make([][]int, w+1)
	for i := range b {
		b[i] = make([]int, n)
	}
	b[0][0] = 1
	for i := 0; i < w; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				t := (a[j] & a[k]) | ((a[j] << 1) & a[k]) | (a[j] & (a[k] << 1))
				if t == 0 {
					b[i+1][k] = (b[i+1][k] + b[i][j]) % MOD
				}
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = (res + b[w][i]) % MOD
	}

	out(res - 1)
}
