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

type pos struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	s := make([]pos, N)
	t := make([]pos, N)
	for i := 0; i < N; i++ {
		s[i] = pos{getI(), getI()}
	}
	for i := 0; i < N; i++ {
		t[i] = pos{getI(), getI()}
	}

	sort.Slice(s, func(i, j int) bool {
		if s[i].x == s[j].x {
			return s[i].y < s[j].y
		}
		return s[i].x < s[j].x
	})
	sort.Slice(t, func(i, j int) bool {
		if t[i].x == t[j].x {
			return t[i].y < t[j].y
		}
		return t[i].x < t[j].x
	})

	// そのまま一致するか？
	ok := true
	for i := 0; i < N; i++ {
		if s[i] != t[i] {
			ok = false
		}
	}
	if ok {
		out("Yes")
		return
	}

	// ｘを反転したものとの差分が一定か？
	p := make([]pos, N)
	for i := 0; i < N; i++ {
		p[i] = pos{-s[i].x, s[i].y}
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})

	ok = true
	diff := p[0].x - t[0].x
	for i := 0; i < N; i++ {
		if p[i].y != t[i].y {
			ok = false
		} else if p[i].x-t[i].x != diff {
			ok = false
		}
	}

	if ok {
		out("Yes")
		return
	}

	// yを反転したものとの差分が一定か？
	p = make([]pos, N)
	for i := 0; i < N; i++ {
		p[i] = pos{s[i].x, -s[i].y}
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})

	ok = true
	diff = p[0].y - t[0].y
	for i := 0; i < N; i++ {
		if p[i].x != t[i].x {
			ok = false
		} else if p[i].y-t[i].y != diff {
			ok = false
		}
	}

	if ok {
		out("Yes")
		return
	}
	out("No")
}
