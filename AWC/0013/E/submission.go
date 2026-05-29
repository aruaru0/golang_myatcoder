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
	N, M := getI(), getI()
	c := make([][]int, N)
	for i := 0; i < N; i++ {
		k := getI()
		c[i] = getInts(k)
	}

	g := NewMFGraph[int](M + N + 2)

	start := 0
	end := M + N + 1
	for i := 0; i < M; i++ {
		g.AddEdge(start, i+1, 1)
	}
	for i := 0; i < N; i++ {
		g.AddEdge(M+i+1, end, 1)
	}
	for i := 0; i < N; i++ {
		for _, e := range c[i] {
			from, to := e, M+i+1
			g.AddEdge(from, to, 1)
		}
	}
	ret := g.Flow(start, end)
	out(ret)
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type MFEdge[Cap Integer] struct {
	From int
	To   int
	Cap  Cap
	Flow Cap
}

type _edge[Cap Integer] struct {
	to  int
	rev int
	cap Cap
}

type MFGraph[Cap Integer] struct {
	n   int
	pos []struct {
		first  int
		second int
	}
	g [][]_edge[Cap]
}

func NewMFGraph[Cap Integer](n int) *MFGraph[Cap] {
	return &MFGraph[Cap]{
		n: n,
		g: make([][]_edge[Cap], n),
	}
}

func (g *MFGraph[Cap]) AddEdge(from, to int, cap Cap) int {
	if from < 0 || from >= g.n {
		panic("from is out of bounds")
	}
	if to < 0 || to >= g.n {
		panic("to is out of bounds")
	}
	if cap < 0 {
		panic("capacity cannot be negative")
	}
	m := len(g.pos)
	g.pos = append(g.pos, struct {
		first  int
		second int
	}{from, len(g.g[from])})
	fromId := len(g.g[from])
	toId := len(g.g[to])
	if from == to {
		toId++
	}
	g.g[from] = append(g.g[from], _edge[Cap]{to, toId, cap})
	g.g[to] = append(g.g[to], _edge[Cap]{from, fromId, 0})
	return m
}

func (g *MFGraph[Cap]) GetEdge(i int) MFEdge[Cap] {
	if i < 0 || i >= len(g.pos) {
		panic("index out of bounds")
	}
	_e := g.g[g.pos[i].first][g.pos[i].second]
	_re := g.g[_e.to][_e.rev]
	return MFEdge[Cap]{
		From: g.pos[i].first,
		To:   _e.to,
		Cap:  _e.cap + _re.cap,
		Flow: _re.cap,
	}
}

func (g *MFGraph[Cap]) Edges() []MFEdge[Cap] {
	m := len(g.pos)
	result := make([]MFEdge[Cap], m)
	for i := 0; i < m; i++ {
		result[i] = g.GetEdge(i)
	}
	return result
}

func (g *MFGraph[Cap]) ChangeEdge(i int, newCap, newFlow Cap) {
	if i < 0 || i >= len(g.pos) {
		panic("index out of bounds")
	}
	if newFlow < 0 || newFlow > newCap {
		panic("invalid flow or capacity")
	}
	first := g.pos[i].first
	second := g.pos[i].second
	_e := &g.g[first][second]
	_re := &g.g[_e.to][_e.rev]
	_e.cap = newCap - newFlow
	_re.cap = newFlow
}

func (g *MFGraph[Cap]) Flow(s, t int) Cap {
	var maxCap Cap = 1
	for {
		if maxCap*2+1 > maxCap {
			maxCap = maxCap*2 + 1
		} else {
			break
		}
	}
	return g.FlowLimit(s, t, maxCap)
}

func maxflow_minCap[Cap Integer](a, b Cap) Cap {
	if a < b {
		return a
	}
	return b
}

func (g *MFGraph[Cap]) FlowLimit(s, t int, flowLimit Cap) Cap {
	if s < 0 || s >= g.n {
		panic("s is out of bounds")
	}
	if t < 0 || t >= g.n {
		panic("t is out of bounds")
	}
	if s == t {
		panic("s and t must be different")
	}

	level := make([]int, g.n)
	iter := make([]int, g.n)
	que := make([]int, 0, g.n)

	bfs := func() {
		for i := range level {
			level[i] = -1
		}
		level[s] = 0
		que = que[:0]
		que = append(que, s)
		for len(que) > 0 {
			v := que[0]
			que = que[1:]
			for _, e := range g.g[v] {
				if e.cap == 0 || level[e.to] >= 0 {
					continue
				}
				level[e.to] = level[v] + 1
				if e.to == t {
					return
				}
				que = append(que, e.to)
			}
		}
	}

	var dfs func(v int, up Cap) Cap
	dfs = func(v int, up Cap) Cap {
		if v == s {
			return up
		}
		res := Cap(0)
		levelV := level[v]
		for ; iter[v] < len(g.g[v]); iter[v]++ {
			i := iter[v]
			e := &g.g[v][i]
			if levelV <= level[e.to] || g.g[e.to][e.rev].cap == 0 {
				continue
			}
			d := dfs(e.to, maxflow_minCap(up-res, g.g[e.to][e.rev].cap))
			if d <= 0 {
				continue
			}
			g.g[v][i].cap += d
			g.g[e.to][e.rev].cap -= d
			res += d
			if res == up {
				return res
			}
		}
		level[v] = g.n
		return res
	}

	flow := Cap(0)
	for flow < flowLimit {
		bfs()
		if level[t] == -1 {
			break
		}
		for i := range iter {
			iter[i] = 0
		}
		f := dfs(t, flowLimit-flow)
		if f == 0 {
			break
		}
		flow += f
	}
	return flow
}

func (g *MFGraph[Cap]) MinCut(s int) []bool {
	visited := make([]bool, g.n)
	que := make([]int, 0, g.n)
	que = append(que, s)
	visited[s] = true
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		for _, e := range g.g[p] {
			if e.cap > 0 && !visited[e.to] {
				visited[e.to] = true
				que = append(que, e.to)
			}
		}
	}
	return visited
}
