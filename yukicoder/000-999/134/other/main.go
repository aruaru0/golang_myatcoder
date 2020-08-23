package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

func getf() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

const inf = math.MaxFloat32

func main() {
	sc.Split(bufio.ScanWords)
	x0, y0 := getInt(), getInt()
	N := getInt() + 1
	x := make([]int, N)
	y := make([]int, N)
	c := make([]float64, N)
	tot := 0.0
	x[0], y[0] = x0, y0
	for i := 1; i < N; i++ {
		x[i], y[i], c[i] = getInt(), getInt(), getf()
		tot += c[i]
	}
	// out(x, y, c)
	n := 1 << N
	dp := make([][]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for mask := 0; mask < n; mask++ {
		for from := 0; from < N; from++ {
			if dp[from][mask] == inf {
				continue
			}
			w := 0.0
			a := make([]int, 0, N)
			for j := 0; j < N; j++ {
				if mask&(1<<j) != 0 {
					w += c[j]
				} else {
					a = append(a, j)
				}
			}
			for _, to := range a {
				l := float64(abs(x[from]-x[to]) + abs(y[from]-y[to]))
				l *= (tot - w + 100.0) / 120.0
				dp[to][mask|(1<<to)] = math.Min(dp[to][mask|(1<<to)], dp[from][mask]+l+c[to])
			}
		}
	}
	// for i := 0; i < N; i++ {
	// 	out(dp[i])
	// }
	out(dp[0][(1<<N)-1])
}
