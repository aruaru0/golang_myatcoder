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

type BIT struct {
	v []int
}

func newBIT(n int) *BIT {
	b := new(BIT)
	b.v = make([]int, n)
	return b
}
func (b BIT) sum(a int) int {
	ret := 0
	for i := a + 1; i > 0; i -= i & -i {
		ret += b.v[i-1]
	}
	return ret
}
func (b BIT) rangeSum(x, y int) int {
	if y == 0 {
		return 0
	}
	y--
	if x == 0 {
		return b.sum(y)
	} else {
		return b.sum(y) - b.sum(x-1)
	}
}
func (b BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
	}
}

func f(v []int) []int {
	m := make(map[int]int)
	for _, e := range v {
		m[e] = 1
	}
	t := make([]int, 0)
	for e := range m {
		t = append(t, e)
	}
	sort.Ints(t)
	for i := 0; i < len(t); i++ {
		m[t[i]] = i
	}
	n := len(v)
	for i := 0; i < n; i++ {
		v[i] = m[v[i]]
	}

	val := make([]int, n)
	bit := newBIT(n + 1)
	for i := 0; i < n; i++ {
		val[i] = bit.rangeSum(v[i]+1, n)
		bit.add(v[i], 1)
	}
	return val
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	c := getInts(N)
	x := getInts(N)

	l := make(map[int][]int)
	p := make(map[int][]int)
	for i := 0; i < N; i++ {
		l[c[i]] = append(l[c[i]], x[i])
		p[c[i]] = append(p[c[i]], i)
	}

	v := make([]int, N)
	for e := range l {
		val := f(l[e])
		for i := 0; i < len(val); i++ {
			v[p[e][i]] = val[i]
		}
	}

	bit := newBIT(N + 2)
	cnt := make([]int, N)
	ans := 0
	for i := 0; i < N; i++ {
		cnt[i] = bit.rangeSum(x[i]+1, N+1)
		bit.add(x[i], 1)
		ans += cnt[i] - v[i]
	}

	out(ans)
}
