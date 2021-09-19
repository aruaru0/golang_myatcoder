package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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

//
// LCAの簡易バージョン（２点の距離計算のついたバージョン）
//

// l := newLCA(n)
// l.addEdge(u,v)
// l.build(root) ※実行が必要
// root := l.lca(u,v)
// dist := l.dist(u,v)

type lca struct {
	n      int
	log    int
	parent [][]int
	dep    []int
	G      [][]int
}

func newLCA(n int) *lca {
	var ret lca
	ret.n = n
	ret.log = int(math.Log2(float64(n))) + 1
	ret.parent = make([][]int, ret.log)
	for i := 0; i < ret.log; i++ {
		ret.parent[i] = make([]int, n)
	}
	ret.dep = make([]int, n)
	ret.G = make([][]int, n)
	return &ret
}

func (l *lca) dfs(v, p, d int) {
	l.parent[0][v] = p
	l.dep[v] = d
	for _, to := range l.G[v] {
		if to == p {
			continue
		}
		l.dfs(to, v, d+1)
	}
}

func (l *lca) addEdge(from, to int) {
	l.G[from] = append(l.G[from], to)
	l.G[to] = append(l.G[to], from)
}

func (l *lca) build(root int) {
	l.dfs(root, -1, 0)
	for k := 0; k+1 < l.log; k++ {
		for v := 0; v < l.n; v++ {
			if l.parent[k][v] < 0 {
				l.parent[k+1][v] = -1
			} else {
				l.parent[k+1][v] = l.parent[k][l.parent[k][v]]
			}
		}
	}
}

func (l *lca) depth(v int) int {
	return l.dep[v]
}

func (l *lca) lca(u, v int) int {
	if l.dep[u] > l.dep[v] {
		u, v = v, u
	}
	for k := 0; k < l.log; k++ {
		if (l.dep[v]-l.dep[u])>>k&1 == 1 {
			v = l.parent[k][v]
		}
	}
	if u == v {
		return u
	}
	for k := l.log - 1; k >= 0; k-- {
		if l.parent[k][u] != l.parent[k][v] {
			u = l.parent[k][u]
			v = l.parent[k][v]
		}
	}
	return l.parent[0][u]
}

func (l *lca) dist(u, v int) int {
	return l.dep[u] + l.dep[v] - 2*l.dep[l.lca(u, v)]
}

type edge struct {
	to, idx int
}

var node [][]edge
var used []bool

func dfs(cur, target, prev int) bool {
	if cur == target {
		return true
	}
	for _, e := range node[cur] {
		if e.to == prev {
			continue
		}
		ret := dfs(e.to, target, cur)
		if ret == true {
			used[e.idx] = true
			return true
		}
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node = make([][]edge, N)
	l := newLCA(N)
	for i := 0; i < N-1; i++ {
		a, b := getI()-1, getI()-1
		node[a] = append(node[a], edge{b, i})
		node[b] = append(node[b], edge{a, i})
		l.addEdge(a, b)
	}
	l.build(0)
	M := getI()
	pat := make([]int, M)
	// M個のパターンで通過する辺をbitで表現する 最大20bit
	for i := 0; i < M; i++ {
		used = make([]bool, N-1)
		u, v := getI()-1, getI()-1
		x := l.lca(u, v)
		dfs(x, u, -1)
		dfs(x, v, -1)
		y := 0
		// out(u, v, used)
		for i := 0; i < N-1; i++ {
			y = y * 2
			if used[i] == false {
				y++
			}
		}
		// out(u, v, l.lca(u, v))
		pat[i] = y
	}

	// Mのパターンを全探索
	n := 1 << M
	ans := 0
	for bit := 0; bit < n; bit++ {
		mask := 1<<(N-1) - 1
		for i := 0; i < M; i++ {
			if (bit>>i)%2 == 1 {
				mask &= pat[i]
			}
		}
		cnt := bits.OnesCount(uint(bit))
		// 包除原理（偶数はプラス、奇数はマイナス）
		if cnt%2 == 0 {
			ans += 1 << bits.OnesCount(uint(mask))
		} else {
			ans -= 1 << bits.OnesCount(uint(mask))
		}
		// fmt.Fprintf(wr, "%4.4b %4.4b cnt %d mask %d\n", bit, mask, cnt, 1<<bits.OnesCount(uint(mask)))
	}
	out(ans)
}
