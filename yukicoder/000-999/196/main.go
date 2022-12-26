package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

const mod = int(1e9 + 7)

var node [][]int

func dfs(v, p int) []int {
	ret := make([]int, 0)
	ret = append(ret, 1)
	for _, e := range node[v] {
		if e == p {
			continue
		}
		t := dfs(e, v)
		n := len(ret)
		m := len(t)
		nret := make([]int, n+m-1)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				nret[i+j] += ret[i] * t[j] % mod
				nret[i+j] %= mod
			}
		}
		ret = nret
	}
	ret = append(ret, 1)
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		a, b := getInt(), getInt()
		node[a] = append(node[a], b)
		node[b] = append(node[b], a)
	}

	ret := dfs(0, -1)[K]

	// for i := 0; i < N; i++ {
	// 	out(i, dp[i])
	// }
	out(ret)
}
