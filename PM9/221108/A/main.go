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

const inf int = 1e18

type edge struct {
	to, cost int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N - 1)
	b := getInts(N - 2)

	node := make([][]edge, N)
	rnode := make([][]edge, N)
	for i := 0; i < N-1; i++ {
		node[i] = append(node[i], edge{i + 1, a[i]})
		rnode[i+1] = append(rnode[i+1], edge{i, a[i]})
	}
	for i := 0; i < N-2; i++ {
		node[i] = append(node[i], edge{i + 2, b[i]})
		rnode[i+2] = append(rnode[i+2], edge{i, b[i]})
	}

	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 0; i < N; i++ {
		for _, e := range node[i] {
			dp[e.to] = min(dp[e.to], dp[i]+e.cost)
		}
	}

	tot := dp[N-1]
	idx := N - 1
	ans := []int{idx}
	for idx != 0 {
		for _, e := range rnode[idx] {
			if tot-e.cost == dp[e.to] {
				tot -= e.cost
				idx = e.to
				ans = append(ans, e.to)
				break
			}
		}
	}

	out(len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(wr, ans[i]+1, " ")
	}
	out()
}
