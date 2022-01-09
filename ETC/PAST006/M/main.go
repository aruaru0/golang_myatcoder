package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	t := redblacktree.NewWithIntComparator()
	const (
		R = iota
		A
	)
	ans := 0
	m := map[int]int{}
	aa := make([][2]int, n)
	for i := 0; i < n; i++ {
		a := getI()
		ans += m[a]
		aa[i][R] = i + 1
		aa[i][A] = a
		t.Put(i, 1)
		m[a]++
	}
	q := getI()
	for i := 0; i < q; i++ {
		l := getI() - 1
		r := getI()
		x := getI()
		c, found := t.Ceiling(l)
		if found {
			k := c.Key.(int)
			for k < r {
				rr := aa[k][R]
				a := aa[k][A]
				ans -= dec(m[a], rr-k)
				m[a] -= rr - k
				t.Remove(k)
				if r < rr {
					t.Put(r, 1)
					aa[r][R] = rr
					aa[r][A] = a
					ans += inc(m[a], rr-r)
					m[a] += (rr - r)
				}
				k = aa[k][R]
			}
		}
		f, found := t.Floor(l - 1)
		if found {
			k := f.Key.(int)
			rr := aa[k][R]
			a := aa[k][A]
			if rr > l {
				ans -= dec(m[a], rr-k)
				m[a] -= (rr - k)
				t.Put(k, 1)
				aa[k][R] = l
				ans += inc(m[a], l-k)
				m[a] += (l - k)
			}
			if rr > r {
				t.Put(r, 1)
				ans += inc(m[a], rr-r)
				m[a] += (rr - r)
				aa[r][R] = rr
				aa[r][A] = a
			}
		}
		t.Put(l, 1)
		ans += inc(m[x], r-l)
		m[x] += r - l
		aa[l][R] = r
		aa[l][A] = x
		out(ans)
	}
}
func inc(x, add int) int {
	return (x*2 + add - 1) * add / 2
}
func dec(x, sub int) int {
	return (x*2 - (sub + 1)) * sub / 2
}
