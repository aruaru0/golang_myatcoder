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

type CaartesianTree struct {
	n, root int
	l, r    []int
}

func newCartesianTree(a []int, _max bool) CaartesianTree {
	ret := CaartesianTree{}
	n := len(a)
	ret.n = n
	ret.l = make([]int, n)
	ret.r = make([]int, n)
	for i := 0; i < n; i++ {
		ret.l[i] = -1
		ret.r[i] = -1
	}

	st := make([]int, 0)

	for i := 0; i < n; i++ {
		p := -1

		for len(st) > 0 {
			topIndex := st[len(st)-1]
			if (a[topIndex] < a[i]) != _max {
				break
			}
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ret.r[j] = p
			p = j
		}

		ret.l[i] = p
		st = append(st, i)
	}

	for i := 0; i < len(st)-1; i++ {
		ret.r[st[i]] = st[i+1]
	}

	if len(st) > 0 {
		ret.root = st[0]
	} else {
		ret.root = -1
	}

	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m, k := getI(), getI(), getI()
	s := getStrings(n)

	ans := 0
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+2)
	}
	tri := func(s int) int {
		if s < 0 {
			return 0
		}
		return (s + 2) * (s + 1) / 2
	}
	for h := 1; h <= n; h++ {
		for c := 0; c < m+2; c++ {
			f[h][c] = f[h-1][c] + tri(k/h-1-c)
		}
	}

	add := func(l, r, h int) {
		l++
		r++
		ans += f[h][0]
		ans -= f[h][l]
		ans -= f[h][r]
		ans += f[h][l+r]
	}

	a := make([]int, m)
	var dfs func(l, r, v int)
	for bi := 0; bi < n; bi++ {
		for i := 0; i < m; i++ {
			a[i]++
			// out(bi, i, s)
			if s[bi][i] == '#' {
				a[i] = 0
			}
		}
		t := newCartesianTree(a, false)
		dfs = func(l, r, v int) {
			if v == -1 {
				return
			}
			add(v-l, r-v, a[v])
			dfs(l, v-1, t.l[v])
			dfs(v+1, r, t.r[v])
		}
		dfs(0, m-1, t.root)
	}

	out(ans)
}
