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

var dp [2005][2005][4]int

func main() {
	// 最初、黒いマスを辿る方法だとミスリーディングしてdpの遷移をミスって考えていた
	// 出題は同じ色を辿ればOK
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	h, w := getI(), getI()
	r := getInts(h)
	c := getInts(w)
	a := make([]string, h)
	for i := 0; i < h; i++ {
		a[i] = getS()
	}

	const inf = int(1e18)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 4; k++ {
				dp[i][j][k] = inf
			}
		}
	}

	// k = 0（反転なし）1（列を反転）2（行を反転）3（両方反転）
	for k := 0; k < 4; k++ {
		co := 0
		if k&1 != 0 {
			co += r[0]
		}
		if k&2 != 0 {
			co += c[0]
		}
		dp[0][0][k] = co
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 4; k++ {
				x := int(a[i][j] - '0')
				// kによって反転するかどうかを決める
				if k&1 != 0 {
					x ^= 1
				}
				if k&2 != 0 {
					x ^= 1
				}
				if i+1 < h {
					// 列方向 nkは次の状態
					nk, y := k&2, int(a[i+1][j]-'0')
					if k&2 != 0 {
						y ^= 1
					}
					co := 0
					if x != y { // 色が違う場合は反転させる
						co += r[i+1]
						nk |= 1
					}
					chmin(&dp[i+1][j][nk], dp[i][j][k]+co)
				}
				if j+1 < w {
					nk, y := k&1, int(a[i][j+1]-'0')
					if k&1 != 0 {
						y ^= 1
					}
					co := 0
					if x != y {
						co += c[j+1]
						nk |= 2
					}
					chmin(&dp[i][j+1][nk], dp[i][j][k]+co)
				}
			}
		}
	}

	ans := inf
	for k := 0; k < 4; k++ {
		ans = min(ans, dp[h-1][w-1][k])
	}
	out(ans)
}
