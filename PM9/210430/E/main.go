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

var H, W, A, B int

func idx2pos(idx int) (int, int) {
	return idx / W, idx % W
}

func pos2idx(y, x int) int {
	return y*W + x
}

var ans = 0

func rec(n, a, b, bit int) {
	// fmt.Fprintf(wr, "%d %d %d %4.4b\n", n, a, b, bit)
	if n == H*W {
		// out("pass")
		ans++
		return
	}

	if (bit>>n)%2 == 1 {
		rec(n+1, a, b, bit)
		return
	}
	if b < B {
		rec(n+1, a, b+1, bit|1<<n)
	}
	if a < A {
		y, x := idx2pos(n)
		if x+1 < W {
			left := pos2idx(y, x+1)
			if (bit>>left)%2 == 0 {
				x := bit | 1<<n
				x |= 1 << left
				rec(n+1, a+1, b, x)
			}
		}
		if y+1 < H {
			down := pos2idx(y+1, x)
			if (bit>>down)%2 == 0 {
				x := bit | 1<<n
				x |= 1 << down
				rec(n+1, a+1, b, x)
			}
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, A, B = getI(), getI(), getI(), getI()

	rec(0, 0, 0, 0)
	out(ans)
}
