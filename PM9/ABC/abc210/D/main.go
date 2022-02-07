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
	H, W, C := getI(), getI(), getI()
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = getInts(W)
	}

	// コストは、問題文の通り、
	//  cost = A_ij + A_i'j' + (|i-i'| + |j-j'|)*C
	//  i > i', j > j'であると仮定すると、
	//  cost = A_ij + A_i'j' + (i-i' + j-j')*C
	//  i,j i'j'を整理すると
	//  A_ij + (i+j)*C  +  A_i'j' - (i'+j')*C
	//  となる。
	//  なので、ijとi'j'は独立に計算できることわがわかる。
	//  左上から計算していけば、i'j'は既に計算しているので、
	//  そこまでの最小値を配列dpに入れながら走査して行けばよい
	ans := inf
	for k := 0; k < 2; k++ {
		// テーブルの初期化。最初、ループ外においてバグらせた
		dp := make([][]int, H)
		for i := 0; i < H; i++ {
			dp[i] = make([]int, W)
			for j := 0; j < W; j++ {
				dp[i][j] = inf
			}
		}
		// 左上から確定させていく
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if i > 0 {
					// 上があるなら、更新
					dp[i][j] = min(dp[i][j], dp[i-1][j])
				}
				if j > 0 {
					// 左があるなら、更新
					dp[i][j] = min(dp[i][j], dp[i][j-1])
				}
				// コストを計算
				ans = min(ans, a[i][j]+(i+j)*C+dp[i][j])
				dp[i][j] = min(dp[i][j], a[i][j]-(i+j)*C)
			}
		}

		// aを上下反転させて２回試す
		for i := 0; i < H/2; i++ {
			a[i], a[H-1-i] = a[H-1-i], a[i]
		}
	}
	out(ans)
}
