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

var used []bool
var node [][]int
var cnt []int

func dfs(cur int) {
	used[cur] = true

	for _, e := range node[cur] {
		if used[e] {
			continue
		}
		cnt[e]--
		if cnt[e] == 1 {
			dfs(e)
		}
	}
}

func dfs2(cur int) {
	used[cur] = true
	for _, e := range node[cur] {
		if used[e] {
			continue
		}
		dfs(e)
	}
}

const mod = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	u, v := make([]int, M), make([]int, M)
	node = make([][]int, N)
	for i := 0; i < M; i++ {
		u[i], v[i] = getI()-1, getI()-1
		node[u[i]] = append(node[u[i]], v[i])
		node[v[i]] = append(node[v[i]], u[i])
	}
	q := make([]int, 0)
	cnt = make([]int, N)
	ok := true
	for i := 0; i < N; i++ {
		cnt[i] = len(node[i])
		if cnt[i] == 0 {
			ok = false
		}
		if cnt[i] == 1 {
			q = append(q, i)
		}
	}

	// out(node)

	// out(q, ok, cnt)

	used = make([]bool, N)
	for _, e := range q {
		dfs(e)
	}

	ok = true
	for i := 0; i < N; i++ {
		if cnt[i] == 0 || cnt[i] > 2 {
			ok = false
		}
	}
	if !ok {
		out("0")
		return
	}

	// out(cnt)
	// out(used)

	node = make([][]int, N)
	for i := 0; i < M; i++ {
		if used[u[i]] || used[v[i]] {
			continue
		}
		node[u[i]] = append(node[u[i]], v[i])
		node[v[i]] = append(node[v[i]], u[i])
	}

	ans := 1
	used = make([]bool, N)
	for i := 0; i < N; i++ {
		if len(node[i]) != 0 && !used[i] {
			ans *= 2
			ans %= mod
			dfs2(i)
		}
	}
	out(ans)
}
