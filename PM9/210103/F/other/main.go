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

// PQ :
type PQ struct {
	value, pos, dir int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	W, H := getI(), getI()
	pq := make([]PQ, 0, H+W)
	for i := 0; i < W; i++ {
		pq = append(pq, PQ{getI(), i, 0})
	}
	for i := 0; i < H; i++ {
		pq = append(pq, PQ{getI(), i, 1})
	}
	sort.Slice(pq, func(i, j int) bool {
		return pq[i].value < pq[j].value
	})

	hbit := newBIT(H + 1)
	wbit := newBIT(W + 1)
	ans := 0
	// out(pq)
	for i := 0; i < len(pq); i++ {
		cur := pq[i]
		if cur.dir == 0 {
			v := hbit.sum(H)
			ans += max(0, H+1-v) * cur.value
			// out("0:", cur, v, ans)
			wbit.add(cur.pos, 1)
		} else {
			v := wbit.sum(W)
			ans += max(0, W+1-v) * cur.value
			// out("1:", cur, v, ans)
			hbit.add(cur.pos, 1)
		}
	}
	out(ans)
}
