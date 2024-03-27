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

type pair struct {
	col, cnt int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, M := getI(), getI(), getI()

	col := make([]int, 210000)
	col[0] = 0

	h := make(map[int]bool)
	w := make(map[int]bool)

	T := make([]int, M)
	A := make([]int, M)
	X := make([]int, M)
	for qi := 0; qi < M; qi++ {
		T[qi], A[qi], X[qi] = getI(), getI()-1, getI()
	}

	for qi := M - 1; qi >= 0; qi-- {
		t, a, x := T[qi], A[qi], X[qi]
		if t == 1 {
			if h[a] == false {
				col[x] += W - len(w)
				h[a] = true
			}
		} else {
			if w[a] == false {
				col[x] += H - len(h)
				w[a] = true
			}
		}
	}

	tot := H * W
	for i := 0; i < 210000; i++ {
		tot -= col[i]
	}
	col[0] += tot
	ans := make([]pair, 0)
	for i := 0; i < 210000; i++ {
		if col[i] != 0 {
			ans = append(ans, pair{i, col[i]})
		}
	}
	out(len(ans))
	for _, e := range ans {
		out(e.col, e.cnt)
	}
}
