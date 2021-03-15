package main

import (
	"bufio"
	"fmt"
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

func (b BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, M := getI(), getI(), getI()

	hmin := make([]int, W+1)
	wmin := make([]int, H+1)
	for i := 0; i < W; i++ {
		hmin[i] = H
	}
	for i := 0; i < H; i++ {
		wmin[i] = W
	}
	for i := 0; i < M; i++ {
		h, w := getI()-1, getI()-1
		hmin[w] = min(hmin[w], h)
		wmin[h] = min(wmin[h], w)
	}

	ans := 0
	m := make([][]int, H+1)
	for i := 0; i < W; i++ {
		ans += hmin[i]
		if hmin[i] == 0 {
			for j := i; j < W; j++ {
				m[hmin[i]] = append(m[hmin[i]], j)
			}
			break
		}
		m[hmin[i]] = append(m[hmin[i]], i)
	}

	// out(ans, m)
	bit := newBIT(W + 1)
	for i := 0; i < H; i++ {
		if wmin[i] == 0 {
			break
		}
		x := bit.sum(wmin[i] - 1)
		// out(x, m[i])
		ans += x
		for _, e := range m[i] {
			bit.add(e, 1)
		}
	}
	out(ans)
}
