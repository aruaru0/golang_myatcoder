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
	num, cnt int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	c := []int{0, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	a := make([]pair, M)
	for i := 0; i < M; i++ {
		x := getI()
		a[i] = pair{x, c[x]}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].num > a[j].num
	})
	dp := make([]int, 11000)
	for i := 0; i < 11000; i++ {
		dp[i] = -int(1e10)
	}
	dp[0] = 0
	for i := 0; i < M; i++ {
		for j := a[i].cnt; j <= 10000; j++ {
			dp[j] = max(dp[j], dp[j-a[i].cnt]+1)
		}
	}
	// out(dp[N])
	// out(a)
	n := dp[N]
	pos := N
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < M; j++ {
			if pos-a[j].cnt < 0 {
				continue
			}
			if dp[pos] == dp[pos-a[j].cnt]+1 {
				ans = append(ans, a[j].num)
				pos -= a[j].cnt
				break
			}
		}
	}

	for _, e := range ans {
		fmt.Fprint(wr, e)
	}
	out()
}
