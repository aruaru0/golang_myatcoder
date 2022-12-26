package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

const mod int = 1e9 + 7

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	h, w, k := getI(), getI(), getI()

	if w == 1 {
		fmt.Println(1)
		return
	}

	dp := [101][10]int{}
	dp[0][1] = 1

	for i := 1; i <= h; i++ {
		// ある高さのつながりをビット表す
		// 左右に番兵の０をつける（ｐの長さはｗ＋２）
		for bit := 0; bit < 1<<(w-1); bit++ {
			p := "0000000000" + strconv.FormatInt(int64(bit), 2) + "0"
			p = p[len(p)-w-1:]
			if strings.Contains(p, "11") {
				continue
			}

			for j := 1; j <= w; j++ {
				pl, pr := j-1, j
				if p[pl] == '0' && p[pr] == '0' {
					// 左右がつながっていない場合は上からそのまま下る
					dp[i][j] += dp[i-1][j]
					dp[i][j] %= mod
					continue
				}

				if p[pl] == '1' {
					// つながっている場合は、双方に加算
					// ←方向への流れは、右の１つ上から来るやつのみ
					dp[i][j-1] += dp[i-1][j]
					dp[i][j-1] %= mod
					// →方向の流れ場、左の１つ上から来るやつのみ
					dp[i][j] += dp[i-1][j-1]
					dp[i][j] %= mod
				}
			}
		}
	}
	out(dp[h][k])
}
