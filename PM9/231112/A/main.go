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

func mul(A, B [][]float64) [][]float64 {
	H := len(A)
	W := len(B[0])
	K := len(A[0])
	C := make([][]float64, W)
	for i := 0; i < W; i++ {
		C[i] = make([]float64, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C
}

func powMatrix(A [][]float64, p int) [][]float64 {
	N := len(A)
	ret := make([][]float64, N)
	for i := 0; i < N; i++ {
		ret[i] = make([]float64, N)
		ret[i][i] = 1
	}

	for p > 0 {
		if p&1 == 1 {
			ret = mul(ret, A)
		}
		A = mul(A, A)
		p >>= 1
	}

	return ret
}

func solve() {
	x, y, z, t := getF(), getF(), getF(), getI()

	a := [][]float64{{1 - x, y, 0}, {0, 1 - y, z}, {x, 0, 1 - z}}

	ret := powMatrix(a, t)
	// out(x, y, z, t, a, ret)

	out(ret[0][0]+ret[0][1]+ret[0][2],
		ret[1][0]+ret[1][1]+ret[1][2],
		ret[2][0]+ret[2][1]+ret[2][2],
	)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	Q := getI()
	for qi := 0; qi < Q; qi++ {
		solve()
		// break
	}
}
