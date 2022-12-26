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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	S := make([]string, 18)
	for i := 0; i < 18; i++ {
		S[i] = getS()
	}

	ok := make([][]int, 1<<12)
	for i := 0; i < 1<<12; i++ {
		ok[i] = make([]int, 64)
	}

	// ok[上２列のパターン][下１列のパターン]が作れるパターンなら１、作れないなら０をあらかじめ作っておく
	for i := 0; i < 1<<12; i++ {
		u := i >> 6
		v := i % 64 * 2
		for w := 0; w < 64; w++ {
			res := 1
			for k := 1; k < 7; k++ {
				x := (u >> (k - 1) & 1) + (w >> (k - 1) & 1) + (v >> (k + 1) & 1) + (v >> k & 1) + (v >> (k - 1) & 1)
				if v>>k&1 != x/3 {
					res = 0
					break
				}
			}
			ok[i][w] = res
		}
	}

	// 行のパターンを作成しておく
	st := make([]string, 64)
	for i := 0; i < 64; i++ {
		st[i] = fmt.Sprintf("%6.6b", i)
	}

	dp := make([]int, 1<<12)
	dp[0] = 1
	for xx := 0; xx < 20; xx++ {
		s := "000000"
		if xx < 18 {
			s = S[xx]
		}
		ndp := make([]int, 1<<12)
		for i := 0; i < 1<<12; i++ {
			v := i % 64
			for w := 0; w < 64; w++ {
				if ok[i][w] == 0 { // 上２行を想定して、作れないパターンであればcontinue
					continue
				}
				flg := true
				for k := 0; k < 6; k++ {
					if st[w][k] != s[k] && s[k] != '?' { // 作れない場合はflgをfalseにする
						flg = false
						break
					}
				}
				if flg {
					ndp[v*64+w] += dp[i] // 作れるパターンの場合は加算
				}
			}
		}
		dp = ndp
	}
	out(dp[0])
}
