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

func solve() {
	n, k := getI(), getI()
	to := make([][]int, n)
	for i := 0; i < n-1; i++ {
		u, v := getI()-1, getI()-1
		to[u] = append(to[u], v)
		to[v] = append(to[v], u)
	}

	c := getInts(k)

	m := 0
	leaf := make([]int, n)
	for i := 0; i < n; i++ {
		if len(to[i]) == 1 {
			leaf[i] = 1
		}
		if leaf[i] == 1 {
			m++
		}
	}

	x := 0
	{
		var f func(v, p int) int
		f = func(v, p int) int {
			res := leaf[v]
			mx := 0
			for _, u := range to[v] {
				if u != p {
					r := f(u, v)
					res += r
					mx = max(mx, r)
				}
			}
			mx = max(mx, m-res)
			if mx*2 <= m {
				x = v
			}
			return res
		}
		f(0, -1)
	}

	usef := make([]bool, n)
	gs := make([][]int, 0)
	for _, v := range to[x] {
		gs = append(gs, []int{})
		var f func(v, p int)
		f = func(v, p int) {
			if leaf[v] == 1 {
				gs[len(gs)-1] = append(gs[len(gs)-1], v)
				usef[v] = true
			}
			for _, u := range to[v] {
				if u != p {
					f(u, v)
				}
			}
		}
		f(v, x)
	}
	sort.Slice(gs, func(i, j int) bool {
		return len(gs[i]) > len(gs[j])
	})

	p := make([]int, 0)
	for _, g := range gs {
		p = append(p, g...)
	}

	o := make([]int, len(p))
	copy(o, p)
	is := make([]int, 0)
	for i := 0; i < 2; i++ {
		for j := 0; j < len(o); j++ {
			if j%2 == i {
				is = append(is, j)
			}
		}
	}
	for i := 0; i < len(o); i++ {
		p[is[i]] = o[i]
	}

	p = append(p, x)
	usef[x] = true

	for i := 0; i < n; i++ {
		if !usef[i] {
			p = append(p, i)
		}
	}

	ci := make([]int, k)
	for i := 0; i < k; i++ {
		ci[i] = i
	}
	sort.Slice(ci, func(i, j int) bool {
		return c[ci[i]] > c[ci[j]]
	})

	pi := 0
	ans := make([]int, n)
	for _, i := range ci {
		if c[i] == 1 && pi < m {
			out("-1")
			return
		}
		for j := 0; j < c[i]; j++ {
			if pi == n {
				break
			}
			ans[p[pi]] = i + 1
			pi++
		}
	}

	outSlice(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		solve()
	}
}
