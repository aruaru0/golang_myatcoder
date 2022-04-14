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

const M = 450

var dp [31][31][31][M + 1]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getS()
	K := getI()
	K = min(K, M)
	n := len(s)

	// K=0, E=1, Y=2に変更してaに格納
	a := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			if s[i] == "KEY"[j] {
				a[i] = j
			}
		}
	}
	// それぞれの文字の位置をpに求める
	p := make([][]int, 3)
	for i := 0; i < n; i++ {
		p[a[i]] = append(p[a[i]], i)
	}

	// K, E, Yの数を求める
	sz := make([]int, 0)
	for i := 0; i < 3; i++ {
		sz = append(sz, len(p[i]))
	}
	// 動的計画法
	dp[0][0][0][0] = 1
	for i := 0; i <= sz[0]; i++ {
		for j := 0; j <= sz[1]; j++ {
			for k := 0; k <= sz[2]; k++ {
				if i < sz[0] { // 使えるKがある場合
					x := 0
					// iより前にあるEの数（転倒数）
					for ni := 0; ni < j; ni++ {
						if p[1][ni] > p[0][i] {
							x++
						}
					}
					// iより前にあるYの数（転倒数）
					for ni := 0; ni < k; ni++ {
						if p[2][ni] > p[0][i] {
							x++
						}
					}
					for t := 0; t < M-x; t++ {
						dp[i+1][j][k][t+x] += dp[i][j][k][t]
					}
				}
				if j < sz[1] { // 使えるEがある場合
					x := 0
					for ni := 0; ni < i; ni++ {
						if p[0][ni] > p[1][j] {
							x++
						}
					}
					for ni := 0; ni < k; ni++ {
						if p[2][ni] > p[1][j] {
							x++
						}
					}
					for t := 0; t < M-x; t++ {
						dp[i][j+1][k][t+x] += dp[i][j][k][t]
					}
				}
				if k < sz[2] { // 使えるYがある場合
					x := 0
					for ni := 0; ni < i; ni++ {
						if p[0][ni] > p[2][k] {
							x++
						}
					}
					for ni := 0; ni < j; ni++ {
						if p[1][ni] > p[2][k] {
							x++
						}
					}
					for t := 0; t < M-x; t++ {
						dp[i][j][k+1][t+x] += dp[i][j][k][t]
					}
				}
			}
		}
	}

	ans := 0
	for t := 0; t < K+1; t++ {
		ans += dp[sz[0]][sz[1]][sz[2]][t]
	}
	out(ans)
}
