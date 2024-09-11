package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
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

type pos struct {
	r, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, Q := getI(), getI(), getI()

	h := make([]*set.Set[int], H)
	for i := 0; i < H; i++ {
		h[i] = set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
		h[i].Insert(-1)
		h[i].Insert(W)
		for j := 0; j < W; j++ {
			h[i].Insert(j)
		}
	}
	w := make([]*set.Set[int], W)
	for i := 0; i < W; i++ {
		w[i] = set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
		w[i].Insert(-1)
		w[i].Insert(H)
		for j := 0; j < H; j++ {
			w[i].Insert(j)
		}
	}

	erase := func(i, j int) {
		h[i].Erase(j)
		w[j].Erase(i)
	}

	// m := make(map[pos]bool)
	for qi := 0; qi < Q; qi++ {
		r, c := getI()-1, getI()-1

		{
			it := h[r].LowerBound(c)
			if it.Value() == c {
				erase(r, c)
				continue
			} else {
				if it.Value() != W {
					erase(r, it.Value())
				}
				it = h[r].LowerBound(c)
				it.Prev()
				if it.Value() != -1 {
					erase(r, it.Value())
				}
			}
		}
		{
			it := w[c].LowerBound(r)
			if it.Value() != H {
				erase(it.Value(), c)
			}
			it = w[c].LowerBound(r)
			it.Prev()
			if it.Value() != -1 {
				erase(it.Value(), c)
			}
		}
	}

	ans := 0
	for i := 0; i < H; i++ {
		ans += h[i].Size() - 2
	}
	out(ans)
}
