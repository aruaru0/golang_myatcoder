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

type animal struct {
	x, y int
	n    string
}

func calc(n [33][33]int) [33][33]int {
	var dp [33][33]int
	dp[0][0] = 1
	for i := 0; i < 33; i++ {
		for j := 0; j < 33; j++ {
			if n[i][j] != 0 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
		// out(dp[i][:10])
	}
	return dp
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	W, H, K, P := getI(), getI(), getI(), getI()
	n := make([][]int, H+1)
	for i := 0; i <= H; i++ {
		n[i] = make([]int, W+1)
	}
	a := make([]animal, K)
	for i := 0; i < K; i++ {
		a[i] = animal{getI(), getI(), getS()}
	}

	N := 1 << K
	ans := -1
	pat := []string{}
	for i := 0; i < N; i++ {
		var n [33][33]int
		sel := make([]string, 0)
		for k := 0; k < K; k++ {
			if (i>>k)&1 == 1 {
				sel = append(sel, a[k].n)
			} else {
				n[a[k].y][a[k].x] = 1
			}
		}
		if len(sel) > P {
			continue
		}
		// out(sel)
		// for j := 0; j <= H; j++ {
		// 	out(n[j][:W+1])
		// }
		ret := calc(n)
		if ans < ret[H][W] {
			ans = ret[H][W]
			pat = sel
			// out(ans, pat)
		}
	}
	out(ans % int(1e9+7))
	for _, e := range pat {
		out(e)
	}
}
