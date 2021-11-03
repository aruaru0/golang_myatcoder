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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	a := getInts(n * 3)

	// dpテーブルの初期化
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -inf
		}
	}
	dp[a[0]][a[1]], dp[a[1]][a[0]] = 0, 0

	// pxの初期化
	px := make([]int, n+1)
	for i := 0; i <= n; i++ {
		px[i] = -inf
	}
	px[a[0]], px[a[1]] = 0, 0

	xx := 0
	add := 0

	for y := 1; y < n; y++ {
		// 先頭の３つの要素を取り出す
		a, b, c := a[3*y-1], a[3*y], a[3*y+1]
		// すべて同じ場合はaddを１つ増やす
		if a == b && b == c {
			add++
			continue
		}
		// パターン列挙
		tmp := make([][3]int, 0)
		for i := 1; i <= n; i++ {
			if b == c { // bとcが同じ場合
				tmp = append(tmp, [3]int{i, a, max(px[i], dp[i][b]+1)})
			} else {
				tmp = append(tmp, [3]int{i, a, px[i]})
			}
			if c == a { // cとaが同じ場合
				tmp = append(tmp, [3]int{i, b, max(px[i], dp[i][c]+1)})
			} else {
				tmp = append(tmp, [3]int{i, b, px[i]})
			}
			if a == b { // aとbが同じ場合
				tmp = append(tmp, [3]int{i, c, max(px[i], dp[i][a]+1)})
			} else {
				tmp = append(tmp, [3]int{i, c, px[i]})
			}
		}

		tmp = append(tmp, [3]int{a, b, max(xx, dp[c][c]+1)})
		tmp = append(tmp, [3]int{b, c, max(xx, dp[a][a]+1)})
		tmp = append(tmp, [3]int{c, a, max(xx, dp[b][b]+1)})

		// 全パターンに対してDPを更新
		for _, e := range tmp {
			p, q, v := e[0], e[1], e[2] // pとqが残っているときの最大値
			dp[p][q] = max(dp[p][q], v)
			px[p] = max(px[p], v)
			dp[q][p] = max(dp[q][p], v)
			px[q] = max(px[q], v)
			xx = max(xx, v)
		}
	}

	xx = max(xx, dp[a[3*n-1]][a[3*n-1]]+1)
	out(xx + add)
}
