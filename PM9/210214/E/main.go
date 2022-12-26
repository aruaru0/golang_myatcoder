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

const mod = int(1e9 + 7)

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

	dp := make([][]int, H)
	cx := make([][]int, H)
	cy := make([][]int, H)
	cxy := make([][]int, H)

	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
		cx[i] = make([]int, W)
		cy[i] = make([]int, W)
		cxy[i] = make([]int, W)
	}

	dp[0][0] = 1
	cx[0][0] = 1
	cy[0][0] = 1
	cxy[0][0] = 1

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if x == 0 && y == 0 {
				continue
			}
			if s[y][x] == '#' {
				continue
			}
			if x > 0 {
				cx[y][x] = cx[y][x-1]
			}
			if y > 0 {
				cy[y][x] = cy[y-1][x]
			}
			if x > 0 && y > 0 {
				cxy[y][x] = cxy[y-1][x-1]
			}
			dp[y][x] = (cx[y][x] + cy[y][x] + cxy[y][x]) % mod
			cx[y][x] += dp[y][x]
			cx[y][x] %= mod
			cy[y][x] += dp[y][x]
			cy[y][x] %= mod
			cxy[y][x] += dp[y][x]
			cxy[y][x] %= mod
		}
	}
	out(dp[H-1][W-1])
}
