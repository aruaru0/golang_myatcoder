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

type SccGraph struct {
	n     int
	edges [][2]int
}
type Csr struct {
	start []int
	elist []int
}
type TwoSat struct {
	n        int
	answer   []bool
	sccGraph *SccGraph
}

func newSccGraph(n int) *SccGraph {
	scc := new(SccGraph)
	scc.n = n
	return scc
}
func (scc *SccGraph) NumVertices() int {
	return scc.n
}
func (scc *SccGraph) AddEdge(from int, to int) {
	scc.edges = append(scc.edges, [2]int{from, to})
}
func (c *Csr) csr(n int, edges [][2]int) {
	c.start = make([]int, n+1)
	c.elist = make([]int, len(edges))
	for _, e := range edges {
		c.start[e[0]+1]++
	}
	for i := 1; i <= n; i++ {
		c.start[i] += c.start[i-1]
	}
	counter := make([]int, n+1)
	copy(counter, c.start)
	for _, e := range edges {
		c.elist[counter[e[0]]] = e[1]
		counter[e[0]]++
	}
}
func (scc *SccGraph) SccIds() (int, []int) {
	g := new(Csr)
	g.csr(scc.n, scc.edges)
	nowOrd, groupNum := 0, 0
	visited, low := make([]int, 0, scc.n), make([]int, scc.n)
	ord, ids := make([]int, scc.n), make([]int, scc.n)
	for i := 0; i < scc.n; i++ {
		ord[i] = -1
	}
	var dfs func(v int)
	dfs = func(v int) {
		low[v], ord[v] = nowOrd, nowOrd
		nowOrd++
		visited = append(visited, v)
		for i := g.start[v]; i < g.start[v+1]; i++ {
			to := g.elist[i]
			if ord[to] == -1 {
				dfs(to)
				low[v] = scc.min(low[v], low[to])
			} else {
				low[v] = scc.min(low[v], ord[to])
			}
		}
		if low[v] == ord[v] {
			for {
				u := visited[len(visited)-1]
				visited = visited[:len(visited)-1]
				ord[u] = scc.n
				ids[u] = groupNum
				if u == v {
					break
				}
			}
			groupNum++
		}
	}
	for i := 0; i < scc.n; i++ {
		if ord[i] == -1 {
			dfs(i)
		}
	}
	for i := 0; i < len(ids); i++ {
		ids[i] = groupNum - 1 - ids[i]
	}
	return groupNum, ids
}
func (scc *SccGraph) min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func (scc *SccGraph) Scc() [][]int {
	groupNum, ids := scc.SccIds()
	counts := make([]int, groupNum)
	for _, x := range ids {
		counts[x]++
	}
	groups := make([][]int, groupNum)
	for i := 0; i < groupNum; i++ {
		groups[i] = make([]int, 0, counts[i])
	}
	for i := 0; i < scc.n; i++ {
		groups[ids[i]] = append(groups[ids[i]], i)
	}
	return groups
}
func newTwoSat(n int) *TwoSat {
	ts := new(TwoSat)
	ts.n = n
	ts.answer = make([]bool, n)
	ts.sccGraph = newSccGraph(n * 2)
	return ts
}
func (ts *TwoSat) AddClause(i int, f bool, j int, g bool) {
	ts.sccGraph.AddEdge(2*i+ts.judge(f, 0, 1), 2*j+ts.judge(g, 1, 0))
	ts.sccGraph.AddEdge(2*j+ts.judge(g, 0, 1), 2*i+ts.judge(f, 1, 0))
}
func (ts *TwoSat) Satisfiable() bool {
	_, id := ts.sccGraph.SccIds()
	for i := 0; i < ts.n; i++ {
		if id[i*2] == id[2*i+1] {
			return false
		}
		ts.answer[i] = id[2*i] < id[2*i+1]
	}
	return true
}
func (ts *TwoSat) judge(f bool, a int, b int) int {
	if f {
		return a
	} else {
		return b
	}
}
func (ts *TwoSat) Answer() []bool {
	return ts.answer
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type pair struct {
	i, j int
}

const MX = 2e6

// 解説を丸写し（Go言語にコンバートのみ）
// two-satっぽいとは考えたが、構成できず
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()

	// 素数を列挙 p
	p := make([]int, 0)
	f := make([]bool, MX+1)
	for i := 2; i <= MX; i++ {
		if f[i] {
			continue
		}
		p = append(p, i)
		for j := i * 2; j <= MX; j += i {
			f[i] = true
		}
	}

	s := len(p)
	// pの逆引きテーブルを作成
	mp := make(map[int]int)
	for i := 0; i < s; i++ {
		mp[p[i]] = i
	}

	// 素因数分解
	var get = func(x int) []int {
		res := make([]int, 0)
		for _, np := range p {
			if np*np > x {
				break
			}
			if x%np != 0 {
				continue
			}
			for x%np == 0 {
				x /= np
			}
			res = append(res, mp[np])
		}
		if x != 1 {
			res = append(res, mp[x])
		}
		return res
	}

	as := make([][]int, n*2)
	vs := make([][]int, s)

	for i := 0; i < n; i++ {
		a, b := getI(), getI()
		// 素因数分解してasに格納
		as[i] = get(a)
		as[n+i] = get(b)
		// 各素数を含むindexをvsに格納
		for _, j := range as[i] {
			vs[j] = append(vs[j], i)
		}
		for _, j := range as[n+i] {
			vs[j] = append(vs[j], n+i)
		}
	}

	// vsのリストの長さの累積和を作成
	ms := make([]int, s+1)
	for i := 0; i < s; i++ {
		ms[i+1] = ms[i] + len(vs[i])
	}

	// 2-sat
	// この構成がいまいちり理解できてない
	ts := newTwoSat(n + ms[s])
	for i := 0; i < s; i++ {
		m := len(vs[i])
		si := n + ms[i]
		for j := 0; j < m-1; j++ {
			ts.AddClause(si+j, false, si+j+1, true)
		}
		for j := 0; j < m; j++ {
			v := vs[i][j]
			val := true
			if v >= n {
				v -= n
				val = false
			}
			ts.AddClause(v, !val, si+j, true)
			if j != 0 {
				ts.AddClause(v, !val, si+j-1, false)
			}
		}
	}
	if ts.Satisfiable() {
		out("Yes")
	} else {
		out("No")
	}
}
