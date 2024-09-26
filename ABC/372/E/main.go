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

type Pair struct {
	col, cnt int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, M := getI(), getI(), getI()

	T, A, X := make([]int, M), make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		T[i], A[i], X[i] = getI(), getI(), getI()
	}

	h := make(map[int]bool)
	w := make(map[int]bool)

	tot := make(map[int]int)
	tot[0] = H * W
	for i := M - 1; i >= 0; i-- {
		t, a, x := T[i], A[i], X[i]

		if t == 1 {
			if h[a] == false {
				h[a] = true
				cnt := W - len(w)
				tot[0] -= cnt
				tot[x] += cnt
			}
		} else {
			if w[a] == false {
				w[a] = true
				cnt := H - len(h)
				tot[0] -= cnt
				tot[x] += cnt
			}
		}
	}

	ans := []Pair{}
	for e := range tot {
		if tot[e] != 0 {
			ans = append(ans, Pair{e, tot[e]})
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].col < ans[j].col
	})

	out(len(ans))
	for _, e := range ans {
		out(e.col, e.cnt)
	}
}
