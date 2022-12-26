package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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

type BV struct {
	n int
	d []uint
}

func newBV(n int) *BV {
	var b BV
	b.n = n
	b.d = make([]uint, n)
	return &b
}

func (b *BV) add(x uint) bool {
	x = b.getMin(x)
	if x == 0 {
		return false
	}
	k := b.n - 1
	for ^x>>k%2 == 1 {
		k--
	}
	for i := 0; i < b.n; i++ {
		if (b.d[i]>>k)%2 == 1 {
			b.d[i] ^= x
		}
	}
	b.d[k] = x
	return true
}

func (b *BV) getMin(x uint) uint {
	for i := 0; i < b.n; i++ {
		if b.d[i] == 0 {
			continue
		}
		if (x>>i)%2 == 1 {
			x ^= b.d[i]
		}
	}
	return x
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	n := 1
	for N != 1<<n {
		n++
	}
	s := make(map[int]bool)
	for i := 0; i < N; i++ {
		s[i] = true
	}
	for i := 0; i < M; i++ {
		a := getI()
		delete(s, a)
	}

	bv := newBV(n)
	b := make([]int, 0)
	for x := range s {
		if bv.add(uint(x)) {
			b = append(b, x)
		}
	}

	if len(b) != n {
		out(-1)
		return
	}

	getBottom := func(x int) int {
		b := x & -x
		return bits.OnesCount(uint(b - 1))
	}

	v := 0
	for i := 1; i < N; i++ {
		u := v ^ b[getBottom(i)]
		out(u, v)
		v = u
	}
}
