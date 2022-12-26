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
	n   int
	bit []int
}

func newBIT(n int) *BIT {
	var ret BIT
	ret.n = n
	ret.bit = make([]int, n+1)
	return &ret
}

func (b *BIT) add(k, v int) {
	k++
	for k <= b.n {
		b.bit[k] += v
		k += k & (-k)
	}
}

func (b *BIT) sum(k int) int {
	if k == 0 {
		return 0
	}
	ret := 0
	for k > 0 {
		ret += b.bit[k]
		k -= k & (-k)
	}
	return ret
}

func (b *BIT) rangeSum(x, y int) int {
	return b.sum(y) - b.sum(x)
}

type pair struct {
	v, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getI() - 1
	}

	bit := newBIT(N + 1)
	cnt := 0
	a := make([]int, 0)
	b := make([]bool, N)
	for i := 0; i < N; i++ {
		x := bit.rangeSum(p[i]+1, N+1)
		// out(p[i], x)
		cnt += x
		bit.add(p[i], 1)
		for j := 0; j < x; j++ {
			pos := i - 1 - j
			// out("pos", pos, i, "-->", i-1-j)
			a = append(a, pos)
			if b[pos] == true {
				out(-1)
				return
			}
			b[pos] = true
		}
	}
	if cnt == N-1 {
		// out(cnt)
		for _, e := range a {
			out(e + 1)
		}
	} else {
		out(-1)
	}
}
