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

type S struct {
	s, g int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, X := getI(), getI()

	// dp[i] := i番目までみた場合の金貨の最大値と銀の消費枚数
	dp := make([]S, X+1)
	for i := 0; i < N; i++ {
		a, b, c := getI(), getI(), getI()
		tmp := make([]S, X+1)
		for j := 0; j < X+1; j++ {
			tmp[j] = dp[j]
		}
		for j := X; j >= 0; j-- {
			if j-b < 0 {
				continue
			}
			if tmp[j].g < dp[j-b].g+c {
				tmp[j].g = dp[j-b].g + c
				tmp[j].s = dp[j-b].s + a
			} else if tmp[j].g == dp[j-b].g+c && tmp[j].s > dp[j-b].s+a {
				tmp[j].s = dp[j-b].s + a
			}
		}
		dp = tmp
	}

	g, s, b := 0, 0, 0
	for i := 0; i < X+1; i++ {
		if dp[i].g > g {
			g = dp[i].g
			s = dp[i].s
			b = X - i
		} else if dp[i].g == g && dp[i].s < s {
			g = dp[i].g
			s = dp[i].s
			b = X - i
		}
	}
	out(g, 1000000000-s, b)
}
