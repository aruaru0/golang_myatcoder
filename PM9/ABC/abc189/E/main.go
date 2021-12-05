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

// 行列xとyの積を計算する
func matmul(x, y [3][3]int) [3][3]int {
	var ret [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				ret[i][k] += x[i][j] * y[j][k]
			}
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
	N := getI()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getI(), getI()
	}
	M := getI()
	dp := make([][3][3]int, M+1)
	dp[0] = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	for i := 0; i < M; i++ {
		op := getI()
		switch op {
		case 1:
			rot90R := [3][3]int{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}}
			dp[i+1] = matmul(rot90R, dp[i])
		case 2:
			rot90L := [3][3]int{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}}
			dp[i+1] = matmul(rot90L, dp[i])
		case 3:
			p := getI()
			shiftX := [3][3]int{{-1, 0, 2 * p}, {0, 1, 0}, {0, 0, 1}}
			dp[i+1] = matmul(shiftX, dp[i])
		case 4:
			p := getI()
			shiftY := [3][3]int{{1, 0, 0}, {0, -1, 2 * p}, {0, 0, 1}}
			dp[i+1] = matmul(shiftY, dp[i])
		}
	}
	Q := getI()
	for i := 0; i < Q; i++ {
		a, b := getI(), getI()-1
		px := dp[a][0][0]*x[b] + dp[a][0][1]*y[b] + dp[a][0][2]
		py := dp[a][1][0]*x[b] + dp[a][1][1]*y[b] + dp[a][1][2]
		out(px, py)
	}
}
