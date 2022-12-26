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
	N := getInt()
	x := make([]int, N)
	y := make([]int, N)
	c := make([]float64, N)
	tot := 0.0
	for i := 0; i < N; i++ {
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
	for i := 0; i < n; i++ {
		if i == 0 {
			w := (tot + 100.0) / 120.0
			for j := 0; j < N; j++ {
				l := abs(x[j]-x0) + abs(y[j]-y0)
				dp[j][i|(1<<j)] = w*float64(l) + c[j]
			}
			continue
		}
		w := 0.0
		a := make([]int, 0, N)
		b := make([]int, 0, N)
		for j := 0; j < N; j++ {
			if i&(1<<j) != 0 {
				w += c[j]
				a = append(a, j)
			} else {
				b = append(b, j)
			}
		}
		for _, t := range b {
			for _, f := range a {
				l := float64(abs(x[f]-x[t]) + abs(y[f]-y[t]))
				l *= (tot - w + 100) / 120.0
				// out(f, t, tot, w, l, c[t])
				dp[t][i|(1<<t)] = math.Min(dp[t][i|(1<<t)], dp[f][i]+c[t]+l)
			}
		}
	}
	// for i := 0; i < N; i++ {
	// 	out(dp[i])
	// }
	ans := inf
	for i := 0; i < N; i++ {
		l := abs(x[i]-x0) + abs(y[i]-y0)
		w := 100 / 120.0
		// out(dp[i][(1<<N)-1] + w*float64(l))
		ans = math.Min(ans, dp[i][(1<<N)-1]+w*float64(l))
	}
	out(ans)
}
