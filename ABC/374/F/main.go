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

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64 |
		~string
}
type CC[T Ordered] struct {
	initialized bool
	xs          []T
}

func newCC[T Ordered]() CC[T] {
	return CC[T]{initialized: false}
}

// struct ccのadd簡数
func (cc *CC[T]) add(x T) {
	cc.xs = append(cc.xs, x)
}

func (cc *CC[T]) init() {
	sort.Slice(cc.xs, func(i, j int) bool {
		return cc.xs[i] < cc.xs[j]
	})
	tmp := make([]T, 0)
	m := make(map[T]bool)
	for _, x := range cc.xs {
		if m[x] {
			continue
		}
		tmp = append(tmp, x)
		m[x] = true
	}
	cc.xs = tmp
	cc.initialized = true
}
func (cc *CC[T]) upperBound(x T) int {
	if cc.initialized == false {
		cc.init()
	}

	idx := sort.Search(len(cc.xs), func(i int) bool {
		return cc.xs[i] > x
	})
	return idx - 1
}

func (cc *CC[T]) get(i int) T {
	if cc.initialized == false {
		cc.init()
	}
	return cc.xs[i]
}

func (cc *CC[T]) len() int {
	if cc.initialized == false {
		cc.init()
	}
	return len(cc.xs)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, k, x := getI(), getI(), getI()
	t := getInts(n)

	const inf = int(1e18)
	cc := newCC[int]()
	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ {
			cc.add(t[i] + x*j)
		}
	}
	cc.add(-1)
	cc.add(inf)

	m := cc.len()
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	order := make([]int, m)
	for i := 0; i < n; i++ {
		order[cc.upperBound(t[i])]++
	}
	orderS := make([]int, m+1)
	for i := 0; i < m; i++ {
		orderS[i+1] = orderS[i] + order[i]
	}
	sum := func(l, r int) int {
		return orderS[r] - orderS[l]
	}

	for i := 0; i < m-1; i++ {
		for j := 0; j <= n; j++ {
			now := dp[i][j]
			if now == inf {
				continue
			}
			{
				nj := j + order[i+1]
				chmin(&dp[i+1][nj], now)
			}
			{
				ni := cc.upperBound(cc.get(i) + x)
				nj := j
				o := min(nj, k)
				nj -= o
				nj += sum(i+1, ni+1)
				chmin(&dp[ni][nj], now+cc.get(i)*o)
			}
		}
	}

	ans := dp[m-1][0]
	for i := 0; i < n; i++ {
		ans -= t[i]
	}
	out(ans)
}
