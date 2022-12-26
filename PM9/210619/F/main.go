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

type binaryIndexedTree []int

func newBinaryIndexedTree(n int) *binaryIndexedTree {
	bit := make(binaryIndexedTree, n+1)
	return &bit
}

func (t binaryIndexedTree) Sum(i int) int {
	s := 0
	for i++; i > 0; i -= i & -i {
		s += t[i]
	}
	return s
}

func (t binaryIndexedTree) RangeSum(l, r int) int {
	return t.Sum(r-1) - t.Sum(l-1)
}

func (t binaryIndexedTree) Add(i, x int) {
	for i++; i < len(t) && i > 0; i += i & -i {
		t[i] += x
	}
}

func (t binaryIndexedTree) LowerBound(x int) int {
	idx, k := 0, 1
	for k < len(t) {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k < len(t) && t[idx+k] < x {
			x -= t[idx+k]
			idx += k
		}
	}
	return idx
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)

	m := make(map[int]int)
	for i := 0; i < N; i++ {
		m[a[i]] = i + 1
	}

	bit := newBinaryIndexedTree(N + 1)
	ans := 0
	for i := 1; i <= N; i++ {
		idx := m[i]
		s := bit.Sum(idx)
		l := bit.LowerBound(s)
		r := bit.LowerBound(s + 1)
		ans += (idx - l) * (r - idx) * i
		bit.Add(idx, 1)
	}
	out(ans)
}
