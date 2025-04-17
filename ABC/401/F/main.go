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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func solveTreeDistances(N int, node [][]int) []int {
	dp1 := make([]int, N) // 各ノードにおける部分木内の最長距離
	dp2 := make([]int, N) // 親から子へ再帰時の距離（全方位木DP）
	ans := make([]int, N)

	// 最初のDFS: 各部分木の最長距離を計算
	var dfs1 func(v, parent int)
	dfs1 = func(v, parent int) {
		for _, to := range node[v] {
			if to == parent {
				continue
			}
			dfs1(to, v)
			if dp1[v] < dp1[to]+1 {
				dp1[v] = dp1[to] + 1
			}
		}
	}

	// 2回目のDFS: 親→子に情報を伝搬（全方位DP）
	var dfs2 func(v, parent int)
	dfs2 = func(v, parent int) {
		// 子供の最大2つの部分木を記録（最長と2番目）
		firstMax, secondMax := -1, -1
		for _, to := range node[v] {
			if to == parent {
				continue
			}
			length := dp1[to] + 1
			if length > firstMax {
				secondMax = firstMax
				firstMax = length
			} else if length > secondMax {
				secondMax = length
			}
		}
		for _, to := range node[v] {
			if to == parent {
				continue
			}
			use := firstMax
			if dp1[to]+1 == firstMax {
				use = secondMax
			}
			dp2[to] = max(dp2[v]+1, use+1)
			dfs2(to, v)
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)

	for i := 0; i < N; i++ {
		ans[i] = max(dp1[i], dp2[i])
	}
	return ans
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N1 := getI()
	node1 := make([][]int, N1)
	for i := 0; i < N1-1; i++ {
		u, v := getI()-1, getI()-1
		node1[u] = append(node1[u], v)
		node1[v] = append(node1[v], u)
	}

	x1 := solveTreeDistances(N1, node1)

	N2 := getI()
	node2 := make([][]int, N2)
	for i := 0; i < N2-1; i++ {
		u, v := getI()-1, getI()-1
		node2[u] = append(node2[u], v)
		node2[v] = append(node2[v], u)
	}

	x2 := solveTreeDistances(N2, node2)

	M := 0
	for _, e := range x1 {
		M = max(M, e)
	}
	for _, e := range x2 {
		M = max(M, e)
	}

	sort.Ints(x2)
	x2sum := make([]int, N2+1)
	for i := 0; i < N2; i++ {
		x2sum[i+1] = x2sum[i] + x2[i]
	}

	tot := 0
	count_total := 0

	for _, a := range x1 {
		threshold_i := (M - 1) - a
		l := upperBound(x2, threshold_i)
		cnt := N2 - l
		count_total += cnt
		sum_b := x2sum[N2] - x2sum[l]
		sum_a_part := a * cnt
		tot += (sum_b + sum_a_part)
	}
	total := tot - (M-1)*count_total
	res := M*N1*N2 + total
	out(res)
}
