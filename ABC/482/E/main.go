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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	node := make([][]int, N)
	for i := 0; i < N-1; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	c := getInts(N)

	var dfs func(cur, prev, cnt int, dist, parent []int)
	dfs = func(cur, prev, cnt int, dist, parent []int) {
		dist[cur] = cnt
		if parent != nil {
			parent[cur] = prev
		}
		for _, e := range node[cur] {
			if e == prev {
				continue
			}
			dfs(e, cur, cnt+1, dist, parent)
		}
	}

	// 1回目のDFS: 適当な頂点(0)から最も遠い頂点Uを見つける
	dist1 := make([]int, N)
	dfs(0, -1, 0, dist1, nil)

	U := 0
	for i := 0; i < N; i++ {
		if dist1[i] > dist1[U] {
			U = i
		}
	}

	// 2回目のDFS: Uから最も遠い頂点Vを見つけつつ、親を記録してパスを復元可能にする
	dist2 := make([]int, N)
	parent := make([]int, N)
	dfs(U, -1, 0, dist2, parent)

	// 直径を求める
	V := 0
	for i := 0; i < N; i++ {
		if dist2[i] > dist2[V] {
			V = i
		}
	}
	L := dist2[V]

	// VからUへ親をたどり、直径のパスを復元する
	path := make([]int, 0)
	cur := V
	for cur != -1 {
		path = append(path, cur)
		cur = parent[cur]
	}

	// out(path)

	ans := 0

	if L%2 == 1 { // 直径が奇数の場合、中心は辺（C1-C2）
		C1 := path[L/2]
		C2 := path[L/2+1]

		count1 := make(map[int]int)
		count2 := make(map[int]int)

		var dfsEdge func(cur, prev, dist int, counter map[int]int, avoid int)
		dfsEdge = func(cur, prev, dist int, counter map[int]int, avoid int) {
			if dist == L/2 {
				counter[c[cur]]++
			}
			for _, e := range node[cur] {
				if e == prev || e == avoid {
					continue
				}
				dfsEdge(e, cur, dist+1, counter, avoid)
			}
		}

		// C1側とC2側へそれぞれ探索（反対側の中心頂点を避ける）
		dfsEdge(C1, -1, 0, count1, C2)
		dfsEdge(C2, -1, 0, count2, C1)

		for col, v1 := range count1 {
			ans += v1 * count2[col]
		}
	} else { // 直径が偶数の場合、中心は1つの頂点(C)
		C := path[L/2]
		totalCount := make(map[int]int)
		subTreeSq := make(map[int]int)

		var dfsNode func(cur, prev, dist int, counter map[int]int)
		dfsNode = func(cur, prev, dist int, counter map[int]int) {
			if dist == L/2 {
				counter[c[cur]]++
			}
			for _, e := range node[cur] {
				if e == prev || e == C {
					continue
				}
				dfsNode(e, cur, dist+1, counter)
			}
		}

		// 中心の隣接頂点（部分木）ごとに独立してカウント
		for _, nxt := range node[C] {
			subCount := make(map[int]int)
			dfsNode(nxt, C, 1, subCount)

			for col, v := range subCount {
				totalCount[col] += v
				subTreeSq[col] += v * v
			}
		}

		// 同じ色の全組み合わせ数から、同じ部分木内の組み合わせ数を引く
		for col, tot := range totalCount {
			ans += (tot*tot - subTreeSq[col]) / 2
		}
	}

	out(ans)
}
