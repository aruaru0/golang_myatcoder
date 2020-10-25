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

var sz []int
var dp []int

const mod = int(1e9 + 7)

func dfs(p int) {
	sz[p] = 1
	for _, e := range node[p] {
		dfs(e)
		sz[p] += sz[e]
		sz[p] %= mod
		dp[p] += dp[e]
		dp[p] %= mod
	}
	dp[p] += sz[p] - 1
	if dp[p] < 0 {
		dp[p] += mod
	}
}

var node [][]int
var cnt int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node = make([][]int, N)
	sub := make([]int, N)
	for i := 0; i < N-1; i++ {
		from, to := getI()-1, getI()-1
		node[from] = append(node[from], to)
		sub[to]++
	}
	top := 0
	for i := 0; i < N; i++ {
		if sub[i] == 0 {
			top = i
		}
	}

	sz = make([]int, N)
	dp = make([]int, N)
	dfs(top)
	ans := 0

	for i := 0; i < N; i++ {
		ans += dp[i]
		ans %= mod
	}
	out(ans)
}
