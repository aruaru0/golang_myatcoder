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

func sel(x, l int) []int {
	ret := make([]int, 0)
	for i := 0; i < l; i++ {
		if (x>>i)%2 == 1 {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H1, W1 := getI(), getI()
	A := make([][]int, H1)
	for i := 0; i < H1; i++ {
		A[i] = getInts(W1)
	}
	H2, W2 := getI(), getI()
	B := make([][]int, H2)
	for i := 0; i < H2; i++ {
		B[i] = getInts(W2)
	}

	check := func(sh, sw []int) bool {
		for i := 0; i < H2; i++ {
			for j := 0; j < W2; j++ {
				sx, sy := sw[j], sh[i]
				// out(B[i][j], A[sy][sx], sy, sx)
				if B[i][j] != A[sy][sx] {
					return false
				}
			}
		}
		return true
	}

	// check([]int{1, 3}, []int{0, 2, 3})
	// return

	h := 1 << H1
	w := 1 << W1
	for i := 0; i < h; i++ {
		selH := sel(i, H1)
		if len(selH) != H2 {
			continue
		}
		for j := 0; j < w; j++ {
			selW := sel(j, W1)
			if len(selW) != W2 {
				continue
			}
			// out(selH, selW)
			if check(selH, selW) {
				out("Yes")
				return
			}
		}
	}
	out("No")
}
