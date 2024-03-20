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

type D struct {
	val, col int
}

type TOP2 struct {
	a, b D
}

const inf = int(1e18)

func newTOP2() TOP2 {
	var ret TOP2
	ret.a = D{-inf, -1}
	ret.b = D{-inf, -1}
	return ret
}

func set(d *TOP2) TOP2 {
	d.a = D{-inf, -1}
	d.b = D{-inf, -1}
	return *d
}

func (t TOP2) add(d D) TOP2 {
	if t.b.val < d.val {
		t.b, d = d, t.b
		if t.a.val < t.b.val {
			t.a, t.b = t.b, t.a
		}
	}
	if t.a.col == t.b.col {
		t.b = d
	}
	return t
}

func (t TOP2) add2(d TOP2) TOP2 {
	t.add(d.a)
	t.add(d.b)
	return t
}

func (t TOP2) get(c int) int {
	if t.a.col == c {
		return t.b.val
	}
	return t.a.val
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, k := getI(), getI()
	c := make([]int, n)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		c[i], v[i] = getI(), getI()
	}

	dp := make([]TOP2, k+1)
	for i := 0; i < k+1; i++ {
		set(&dp[i])
	}
	dp[0] = dp[0].add(D{0, -1})

	for i := 0; i < n; i++ {
		tmp := make([]TOP2, k+1)
		for j := 0; j < k+1; j++ {
			set(&tmp[j])
		}
		for j := 0; j < k+1; j++ {
			if j < k {
				tmp[j+1] = dp[j]
			}
			tmp[j] = tmp[j].add(D{dp[j].get(c[i]) + v[i], c[i]})
		}
		dp = tmp
	}

	ans := dp[k].a.val
	if ans < 0 {
		ans = -1
	}
	out(ans)
	return
}
