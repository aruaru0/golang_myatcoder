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

// 写経（解けてない）

var n, x int
var y, z []int
var dp [3005][3005][2]int
var ok [3005][3005][2]bool
var cd [3005]int
var kind [3005]int

func solve(l, r, sid int) int {
	if ok[l][r][sid] {
		return dp[l][r][sid]
	}
	if sid == 0 {
		if cd[l] != 0 && kind[l] == 0 {
			return 0
		}
		if l != 0 {
			jud := true
			if kind[l-1] > 0 {
				tg := kind[l-1] - 1
				if !(cd[l] <= z[tg] && z[tg] <= cd[r]) {
					jud = false
				}
			}
			if jud {
				dis := cd[l] - cd[l-1]
				dp[l][r][sid] = min(dp[l][r][sid], dis+solve(l-1, r, 0))
			}
		}
		if r != (2*n + 1) {
			jud := true
			if kind[r+1] > 0 {
				tg := kind[r+1] - 1
				if !(cd[l] <= z[tg] && z[tg] <= cd[r]) {
					jud = false
				}
			}

			if jud {
				dis := cd[r+1] - cd[l]
				dp[l][r][sid] = min(dp[l][r][sid], dis+solve(l, r+1, 1))
			}
		}
	} else {
		if cd[r] != 0 && kind[r] == 0 {
			return 0
		}

		if l != 0 {
			jud := true
			if kind[l-1] > 0 {
				tg := kind[l-1] - 1
				if !(cd[l] <= z[tg] && z[tg] <= cd[r]) {
					jud = false
				}
			}

			if jud {
				dis := cd[r] - cd[l-1]
				dp[l][r][sid] = min(dp[l][r][sid], dis+solve(l-1, r, 0))
			}
		}
		if r != (2*n + 1) {
			jud := true
			if kind[r+1] > 0 {
				tg := kind[r+1] - 1
				if !(cd[l] <= z[tg] && z[tg] <= cd[r]) {
					jud = false
				}
			}

			if jud {
				dis := cd[r+1] - cd[r]
				dp[l][r][sid] = min(dp[l][r][sid], dis+solve(l, r+1, 1))
			}
		}
	}

	ok[l][r][sid] = true
	return dp[l][r][sid]
}

type pair struct {
	a, b int
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	n, x = getI(), getI()
	y = getInts(n)
	z = getInts(n)

	for i := 0; i < 3005; i++ {
		for j := 0; j < 3005; j++ {
			dp[i][j][0] = inf
			dp[i][j][1] = inf
		}
	}

	pv := make([]pair, 0)
	pv = append(pv, pair{x, 0})
	pv = append(pv, pair{0, 0})
	// 正が壁、負がハンマー
	for i := 0; i < n; i++ {
		pv = append(pv, pair{y[i], (i + 1)})
		pv = append(pv, pair{z[i], -(i + 1)})
	}

	// 座標でソート
	sort.Slice(pv, func(i, j int) bool {
		if pv[i].a == pv[j].a {
			return pv[i].b < pv[j].b
		}
		return pv[i].a < pv[j].a
	})

	sta := 0
	for i := 0; i < 2*n+2; i++ {
		cd[i] = pv[i].a
		kind[i] = pv[i].b
		if cd[i] == 0 {
			sta = i
		}
	}

	res := solve(sta, sta, 0)
	if res == inf {
		res = -1
	}
	out(res)
}
