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

const off = 15000

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = getInts(W)
	}
	b := make([][]int, H)
	for i := 0; i < H; i++ {
		b[i] = getInts(W)
	}

	dp := make([][][off * 2]bool, H)
	for i := 0; i < H; i++ {
		dp[i] = make([][off * 2]bool, W)
	}
	dp[0][0][off+b[0][0]-a[0][0]] = true
	dp[0][0][off+a[0][0]-b[0][0]] = true
	dx := []int{-1, 0}
	dy := []int{0, -1}
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			for i := 0; i < 2; i++ {
				px := x + dx[i]
				py := y + dy[i]
				if px < 0 || py < 0 {
					continue
				}
				for j := 0; j < off*2; j++ {
					if dp[py][px][j] == true {
						dp[y][x][j+b[y][x]-a[y][x]] = true
						dp[y][x][j-b[y][x]+a[y][x]] = true
					}
				}
			}
			// out(dp[y][x])
		}
	}

	ans := int(1e18)
	for i, e := range dp[H-1][W-1] {
		if e == true {
			ans = min(ans, abs(i-off))
		}
	}
	out(ans)
}
