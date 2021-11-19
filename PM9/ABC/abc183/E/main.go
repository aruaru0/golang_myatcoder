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

const mod = int(1e9) + 7

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H)

	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	dp := make([][][4]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([][4]int, W)
	}

	dp[0][0][0] = 1
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			// 横に進む
			if x != 0 {
				if s[y][x] != '#' {
					dp[y][x][0] += dp[y][x-1][0] + dp[y][x-1][1]
					dp[y][x][1] += dp[y][x-1][0] + dp[y][x-1][1]
				}
			}
			// 下に進む
			if y != 0 {
				if s[y][x] != '#' {
					dp[y][x][0] += dp[y-1][x][0] + dp[y-1][x][2]
					dp[y][x][2] += dp[y-1][x][0] + dp[y-1][x][2]
				}
			}
			// 斜め下に進む
			if x != 0 && y != 0 {
				if s[y][x] != '#' {
					dp[y][x][0] += dp[y-1][x-1][0] + dp[y-1][x-1][3]
					dp[y][x][3] += dp[y-1][x-1][0] + dp[y-1][x-1][3]
				}
			}

			for i := 0; i < 4; i++ {
				dp[y][x][i] %= mod
			}
		}
		// out(dp[y])
	}
	out(dp[H-1][W-1][0])
}
