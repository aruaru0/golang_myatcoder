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
	if len(s) == 0 {
		return
	}
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getI(), getI()
	}

	x = append(x, x...)
	y = append(y, y...)
	x = append(x, x[0])
	y = append(y, y[0])

	n := len(x)
	a := make([]int, n)
	cx := make([]int, n)
	cy := make([]int, n)
	for i := 1; i < n; i++ {
		a[i] = (x[i-1]*y[i] - x[i]*y[i-1])
		cx[i] = (x[i-1] + x[i]) * (x[i-1]*y[i] - x[i]*y[i-1])
		cy[i] = (y[i-1] + y[i]) * (x[i-1]*y[i] - x[i]*y[i-1])
	}

	for i := 1; i < n; i++ {
		a[i] += a[i-1]
		cx[i] += cx[i-1]
		cy[i] += cy[i-1]
	}

	for qi := 0; qi < Q; qi++ {
		u, v := getI(), getI()
		if u > v {
			v += N
		}

		u_idx := u - 1
		v_idx := v - 1

		chord_a := x[v_idx]*y[u_idx] - x[u_idx]*y[v_idx]
		chord_cx := (x[v_idx] + x[u_idx]) * chord_a
		chord_cy := (y[v_idx] + y[u_idx]) * chord_a

		adiff := float64(a[v-1] - a[u-1] + chord_a)
		cxdiff := float64(cx[v-1] - cx[u-1] + chord_cx)
		cydiff := float64(cy[v-1] - cy[u-1] + chord_cy)

		cx_pos := cxdiff / (3.0 * adiff)
		cy_pos := cydiff / (3.0 * adiff)
		out(cx_pos, cy_pos)
	}

}
