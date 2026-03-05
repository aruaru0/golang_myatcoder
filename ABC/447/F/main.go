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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

func solve() {
	N := getI()

	G := make([][]int, N)
	D := make([]int, N)

	for i := 0; i < N-1; i++ {
		u, v := getI()-1, getI()-1
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
		D[u]++
		D[v]++
	}

	// BFSを用いて葉から根へ向かう順序（order）を取得
	order := make([]int, 0, N)
	parent := make([]int, N)
	visited := make([]bool, N)
	q := make([]int, 0, N)

	// 根を頂点0としてBFSを開始
	q = append(q, 0)
	visited[0] = true

	for i := 0; i < len(q); i++ {
		u := q[i]
		order = append(order, u)
		for _, v := range G[u] {
			if !visited[v] {
				visited[v] = true
				parent[v] = u
				q = append(q, v)
			}
		}
	}

	// DP用配列の初期化
	const INF = 1000000
	f := make([]int, N)
	E := make([]int, N)
	for i := 0; i < N; i++ {
		f[i] = -INF
		E[i] = -INF
	}

	ans := 0

	// ボトムアップに木DPを計算
	for i := len(order) - 1; i >= 0; i-- {
		u := order[i]
		max_E1, max_E2 := -INF, -INF

		// 子ノード方向へ伸ばせる最大長の上位2つを取得
		for _, v := range G[u] {
			if v == parent[u] {
				continue
			}
			e := E[v]
			if e > max_E1 {
				max_E2 = max_E1
				max_E1 = e
			} else if e > max_E2 {
				max_E2 = e
			}
		}

		// 自身が中間点になれる場合
		if D[u] >= 4 {
			f[u] = 1 + max_E1
		} else {
			f[u] = -INF
		}

		// 自身が端点になれる場合
		if D[u] >= 3 {
			E[u] = max(f[u], 1) // テンプレートのmax関数を利用
		} else {
			E[u] = -INF
		}

		// uをパスの最上位(LCA)とした場合の最大長を更新
		if D[u] >= 2 {
			if 1 > ans {
				ans = 1
			}
		}
		if D[u] >= 3 {
			val := 1 + max_E1
			if val > ans {
				ans = val
			}
		}
		if D[u] >= 4 {
			val := 1 + max_E1 + max_E2
			if val > ans {
				ans = val
			}
		}
	}

	// 結果の出力
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	Q := getI()
	for i := 0; i < Q; i++ {
		solve()
	}
}
