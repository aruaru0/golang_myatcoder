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

type pos struct {
	x, y int
	d    [3]int
}

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	l := make([]int, 2100)
	l[0] = 1
	for i := 1; i < 2100; i++ {
		l[i] = l[i-1] * 2
		l[i] %= mod
	}

	dp := make([][]int, H)
	x := make([][]int, H)
	y := make([][]int, H)
	z := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
		x[i] = make([]int, W)
		y[i] = make([]int, W)
		z[i] = make([]int, W)
	}

	dp[0][0] = 1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if s[i][j] == '#' {
				continue
			}
			if j > 0 {
				x[i][j] = (x[i][j-1] + dp[i][j-1]) % mod
			}
			if i > 0 {
				y[i][j] = (y[i-1][j] + dp[i-1][j]) % mod
			}
			if i > 0 && j > 0 {
				z[i][j] = (z[i-1][j-1] + dp[i-1][j-1]) % mod
			}
			dp[i][j] = (x[i][j] + y[i][j] + z[i][j]) % mod
		}
	}
	out(dp[H-1][W-1])
}
