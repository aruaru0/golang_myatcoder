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

type pos struct {
	x, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, S, Q := getI(), getI()-1, getI()
	x := make([]pos, N)
	for i := 0; i < N; i++ {
		x[i] = pos{getI(), i}
	}
	const inf = int(1e18)
	sort.Slice(x, func(i, j int) bool {
		return x[i].x < x[j].x
	})

	p := make([][]int, 65)
	for i := 0; i < 65; i++ {
		p[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		if i == 0 {
			p[0][i] = 1
			continue
		}
		if i == N-1 {
			p[0][i] = N - 2
			continue
		}

		li, l := x[i-1].idx, x[i-1].x
		ri, r := x[i+1].idx, x[i+1].x
		c := x[i].x

		// out(li, ri, ":", l, c, r, ":", abs(c-l), abs(c-r))

		if abs(c-l) == abs(c-r) {
			if li < ri {
				p[0][i] = i - 1
			} else {
				p[0][i] = i + 1
			}
		} else if abs(c-l) < abs(c-r) {
			p[0][i] = i - 1
		} else {
			p[0][i] = i + 1
		}
	}

	for i := 1; i < 65; i++ {
		for j := 0; j < N; j++ {
			p[i][j] = p[i-1][p[i-1][j]]
		}
	}

	cur := 0
	for i := 0; i < N; i++ {
		if x[i].idx == S {
			cur = i
			break
		}
	}

	for i := 0; Q > 0; i++ {
		if Q%2 == 1 {
			cur = p[i][cur]
		}
		Q >>= 1
	}

	// out(cur, x[cur])
	out(x[cur].idx + 1)
}
