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

// CartesianTree 構造体
type CartesianTree struct {
	N    int
	Root int
	L    []int
	R    []int
}

func NewCartesianTree(a []int, maxTree bool) *CartesianTree {
	n := len(a)
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = -1
		r[i] = -1
	}

	st := make([]int, 0)

	for i := 0; i < n; i++ {
		p := -1

		for len(st) > 0 {
			top := st[len(st)-1]

			shouldPop := false
			if maxTree {
				if a[top] < a[i] {
					shouldPop = true
				}
			} else {
				if a[top] > a[i] {
					shouldPop = true
				}
			}

			if !shouldPop {
				break
			}

			st = st[:len(st)-1]
			r[top] = p
			p = top
		}

		l[i] = p
		st = append(st, i)
	}

	for i := 0; i < len(st)-1; i++ {
		r[st[i]] = st[i+1]
	}

	root := 0
	if len(st) > 0 {
		root = st[0]
	}

	return &CartesianTree{
		N:    n,
		Root: root,
		L:    l,
		R:    r,
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	p := getInts(N)

	t := NewCartesianTree(p, true)

	out(t)

	var f func(v int) int
	f = func(v int) int {
		l, r := t.L[v], t.R[v]
		res := 0
		if l != -1 {
			res = max(res, f(l)+(v-l))
		}
		if r != -1 {
			res = max(res, f(r)+(r-v))
		}
		return res
	}

	out(f(t.Root))
}
