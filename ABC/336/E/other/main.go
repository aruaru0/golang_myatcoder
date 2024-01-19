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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getS()
	ans := 0
	for s := 1; s <= 9*14; s++ {
		dp := make([][140][9 * 14][2]int, len(N)+1)
		dp[0][0][0][1] = 1
		for d := 0; d < len(N); d++ {
			for i := 0; i < s+1; i++ {
				for j := 0; j < s; j++ {
					for f := 0; f < 2; f++ {
						for t := 0; t < 10; t++ {
							if i+t > s {
								continue
							}
							// 一番上の桁で、tが桁より大きい場合は処理しない
							if f != 0 && (int(N[d]-'0')) < t {
								continue
							}
							// 一番上の桁が、tと一致しているかどうかのフラグを作成
							flg := 0
							if int(N[d]-'0') == t && f != 0 {
								flg = 1
							}
							dp[d+1][i+t][(j*10+t)%s][flg] += dp[d][i][j][f]
						}
					}
				}
			}
		}
		ans += dp[len(N)][s][0][0] + dp[len(N)][s][0][1]
	}
	out(ans)
}
