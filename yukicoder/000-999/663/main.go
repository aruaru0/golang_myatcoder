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

func g(a, b, c int) int {
	if b == 0 && c == 0 {
		return 0
	}
	if a == 1 && b == 1 && c == 1 {
		return 0
	}
	return 1
}

func solve(v0, vn int) int {
	dp := make([][2][2]int, N)

	// v0 ......... Vnのパターンで...の部分をＤＰする
	dp[0][vn][v0] = 1
	for i := 1; i < N-1; i++ {
		for a := 0; a < 2; a++ {
			for b := 0; b < 2; b++ {
				for c := 0; c < 2; c++ {
					if g(a, b, c) == e[i] {
						dp[i][b][c] += dp[i-1][a][b]
						dp[i][b][c] %= mod
					}
				}
			}
		}
	}

	//　先頭と末尾のつじつまがあう場合のみ加算する
	ret := 0
	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			if g(a, b, vn) == e[N-1] && g(b, vn, v0) == e[0] {
				// out(a, b, dp[N-1][a][b])
				ret += dp[N-2][a][b]
				ret %= mod
			}
		}
	}
	return ret
}

var N int
var e []int

const mod = int(1e9 + 7)

func main() {

	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	N = getI()
	e = getInts(N)

	// 先頭と末尾の全パターンを検索してみる
	ans := 0
	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			ans += solve(a, b)
			ans %= mod
		}
	}
	out(ans)

}
