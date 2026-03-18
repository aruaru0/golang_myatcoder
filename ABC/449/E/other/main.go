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

// (累積和が k+1 以上になる最初のインデックス)
func (b *BIT) getKth(k int) int {
	n := len(b.v)
	x := 0
	w := 1

	// n以下の最大の2の冪乗を求める
	for w*2 <= n {
		w *= 2
	}

	for w > 0 {
		if x+w <= n && b.v[x+w-1] < k+1 {
			k -= b.v[x+w-1]
			x += w
		}
		w /= 2
	}
	return x
}

type QS struct {
	pos, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = getI() - 1
	}
	q := getI()
	ans := make([]int, q)
	qs := make([]QS, 0)
	for qi := 0; qi < q; qi++ {
		x := getI() - 1
		if x < n {
			ans[qi] = a[x]
		} else {
			qs = append(qs, QS{x, qi})
		}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].pos > qs[j].pos
	})

	adds := make([][]int, n+1)
	{
		cnt := make([]int, m)
		for i := 0; i < n; i++ {
			cnt[a[i]]++
		}
		for i := 0; i < m; i++ {
			adds[cnt[i]] = append(adds[cnt[i]], i)
		}
	}

	bit := newBIT(m)
	num := 0
	l := n
	for c := 0; c < n+1; c++ {
		for _, i := range adds[c] {
			bit.add(i, 1)
			num++
		}
		for len(qs) != 0 && qs[len(qs)-1].pos < l+num {
			x, qi := qs[len(qs)-1].pos, qs[len(qs)-1].idx
			qs = qs[:len(qs)-1]
			ans[qi] = bit.getKth(x - l)
		}
		l += num
	}

	for len(qs) != 0 {
		x, qi := qs[len(qs)-1].pos, qs[len(qs)-1].idx
		qs = qs[:len(qs)-1]
		ans[qi] = (x - l) % m
	}

	for i := 0; i < q; i++ {
		out(ans[i] + 1)
	}

}
