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

type RARQ struct {
	b1, b2 []int // b1: D[j], b2: j*D[j]
	n      int
}

func newRARQ(n int) *RARQ {
	return &RARQ{b1: make([]int, n+1), b2: make([]int, n+1), n: n}
}

func (t *RARQ) add(bit []int, i, w int) {
	for i++; i <= t.n; i += i & -i {
		bit[i] += w
	}
}
func (t *RARQ) sum(bit []int, i int) int {
	ret := 0
	for i++; i > 0; i -= i & -i {
		ret += bit[i]
	}
	return ret
}

func (t *RARQ) rangeAdd(l, r, v int) {
	t.add(t.b1, l, v)
	t.add(t.b2, l, v*l)
	if r+1 < t.n {
		t.add(t.b1, r+1, -v)
		t.add(t.b2, r+1, -v*(r+1))
	}
}

func (t *RARQ) prefix(p int) int {
	if p < 0 {
		return 0
	}
	return (p+1)*t.sum(t.b1, p) - t.sum(t.b2, p)
}

func (t *RARQ) rangeSum(x, y int) int {
	return t.prefix(y-1) - t.prefix(x-1)
}

type tuple struct {
	l, r, qi int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m, q := getI(), getI(), getI()

	L := make([]int, n)
	R := make([]int, n)
	for i := 0; i < n; i++ {
		L[i], R[i] = getI()-1, getI()
	}

	qs := make([][]tuple, n)
	for qi := 0; qi < q; qi++ {
		u, d, l, r := getI()-2, getI()-1, getI()-1, getI()
		if u >= 0 {
			qs[u] = append(qs[u], tuple{l, r, q + qi})
		}
		qs[d] = append(qs[d], tuple{l, r, qi})
	}

	st := newRARQ(m) // newBIT(m) の代わり

	ans := make([]int, q)
	for i := 0; i < n; i++ {
		st.rangeAdd(L[i], R[i]-1, 1) // 0-indexed [L_i-1, R_i-1] に +1

		for _, e := range qs[i] {
			sign := 1
			l, r, qi := e.l, e.r, e.qi
			if qi >= q {
				qi -= q
				sign = -1
			}
			ans[qi] += st.rangeSum(l, r) * sign
		}
	}

	for i := 0; i < q; i++ {
		out(ans[i])
	}
}
