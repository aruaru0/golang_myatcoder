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

var dp [101][101][101]float32

func rec(a, b, c, x int) float32 {
	n := float32(a + b + c)

	if a == 100 {
		return float32(x - 1)
	}
	if b == 100 {
		return float32(x - 1)
	}
	if c == 100 {
		return float32(x - 1)
	}

	if dp[a][b][c] != 0 {
		return dp[a][b][c]
	}

	ret := float32(0.0)
	if a < 100 {
		r := rec(a+1, b, c, x+1)
		ret += r * (float32(a) / n)
	}
	if b < 100 {
		r := rec(a, b+1, c, x+1)
		ret += r * (float32(b) / n)
	}
	if c < 100 {
		r := rec(a, b, c+1, x+1)
		ret += r * (float32(c) / n)
	}
	dp[a][b][c] = ret
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{}, 10000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	A, B, C := getI(), getI(), getI()

	ret := rec(A, B, C, 1)
	out(ret)
}
