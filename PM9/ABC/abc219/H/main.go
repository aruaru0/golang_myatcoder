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

type pair struct {
	x, a int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	// 負側と正側を分けて管理(0,0)は初期値
	ls := make([]pair, 1)
	rs := make([]pair, 1)
	for i := 0; i < n; i++ {
		x, a := getI(), getI()
		if x < 0 {
			// l側は符号を反転して管理
			ls = append(ls, pair{-x, a})
		} else {
			rs = append(rs, pair{x, a})
		}
	}
	ln := len(ls) - 1
	rn := len(rs) - 1
	// 距離順にソート
	sort.Slice(ls, func(i, j int) bool {
		if ls[i].x == ls[j].x {
			return ls[i].a < ls[j].a
		}
		return ls[i].x < ls[j].x
	})
	sort.Slice(rs, func(i, j int) bool {
		if rs[i].x == rs[j].x {
			return rs[i].a < rs[j].a
		}
		return rs[i].x < rs[j].x
	})

	// DPテーブルの初期化
	// dp[左右][l][r][k]
	const inf = int(1e18)
	dp := make([][][][2]int, ln+1)
	for i := 0; i <= ln; i++ {
		dp[i] = make([][][2]int, rn+1)
		for j := 0; j <= rn; j++ {
			dp[i][j] = make([][2]int, n+1)
			for k := 0; k <= n; k++ {
				dp[i][j][k][0] = -inf
				dp[i][j][k][1] = -inf
			}
		}
	}
	// 初期値を設定
	for s := 0; s < 2; s++ {
		for k := 0; k < n+1; k++ {
			dp[0][0][k][s] = 0
		}
	}
	for i := 0; i < ln+1; i++ {
		for j := 0; j < rn+1; j++ {
			for k := 0; k < n+1; k++ {
				for s := 0; s < 2; s++ {
					for ns := 0; ns < 2; ns++ {
						ni, nj := i, j
						if ns == 0 {
							ni++
						} else {
							nj++
						}
						if ni > ln {
							continue
						}
						if nj > rn {
							continue
						}
						d, a := 0, 0
						if s != ns {
							d = ls[i].x + rs[j].x
						}
						if ns == 0 {
							d += ls[ni].x - ls[i].x
						} else {
							d += rs[nj].x - rs[j].x
						}
						if ns == 0 {
							a = ls[ni].a
						} else {
							a = rs[nj].a
						}
						now := dp[i][j][k][s] - d*k
						chmax(&dp[ni][nj][k][ns], now)
						if k != 0 {
							chmax(&dp[ni][nj][k-1][ns], now+a)
						}
					}
				}
			}
		}
	}
	ans := -inf
	for s := 0; s < 2; s++ {
		ans = max(ans, dp[ln][rn][0][s])
	}
	out(ans)
}
