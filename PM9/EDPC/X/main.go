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

type box struct {
	w, s, v int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	b := make([]box, N)
	for i := 0; i < N; i++ {
		b[i] = box{getI(), getI(), getI()}
	}

	// min(S[a], S[b]-W[a]) ＞ min(S[b], S[a]-W[b])
	sort.Slice(b, func(i, j int) bool {
		return min(b[i].s, b[j].s-b[i].w) > min(b[j].s, b[i].s-b[j].w)
	})

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, 31000)
	}

	for i := 0; i < N; i++ {
		for tot := 0; tot < 21000; tot++ {
			chmax(&dp[i+1][tot], dp[i][tot])
			if tot <= b[i].s {
				chmax(&dp[i+1][tot+b[i].w], dp[i][tot]+b[i].v)
			}
		}
	}

	ans := 0
	for i := 0; i < 21000; i++ {
		ans = max(ans, dp[N][i])
	}
	out(ans)
}
