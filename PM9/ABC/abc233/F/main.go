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

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	n            int
}

// newDsu :
func newDsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
	}
	return &d
}

// Merge :
func (d DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

// Same :
func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader :
func (d DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

// Size :
func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

// Groups : original implement
func (d DSU) Groups() [][]int {
	m := make(map[int][]int)
	for i := 0; i < d.n; i++ {
		x := d.Leader(i)
		if x < 0 {
			m[i] = append(m[i], i)
		} else {
			m[x] = append(m[x], i)
		}
	}
	ret := make([][]int, len(m))
	idx := 0
	for _, e := range m {
		ret[idx] = make([]int, len(e))
		copy(ret[idx], e)
		idx++
	}
	return ret
}

type edge struct {
	to, idx int
}

var p []int
var node [][]edge
var cnt []int
var ans []int

func dfs(cur, prev, target int) bool {
	// out("dfs", cur, target)
	if p[cur] == target {
		return true
	}
	for _, e := range node[cur] {
		if e.to == prev {
			continue
		}
		if cnt[e.to] == 0 {
			continue
		}
		ret := dfs(e.to, cur, target)
		if ret {
			// out("swap", cur, e.to)
			ans = append(ans, e.idx+1)
			p[cur], p[e.to] = p[e.to], p[cur]
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
	uf := newDsu(N)
	p = getInts(N)
	M := getI()
	node = make([][]edge, N)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		if !uf.Same(a, b) {
			uf.Merge(a, b)
			node[a] = append(node[a], edge{b, i})
			node[b] = append(node[b], edge{a, i})
		}
	}

	for i := 0; i < N; i++ {
		p[i]--
		if !uf.Same(i, p[i]) {
			out(-1)
			return
		}
	}

	q := make([]int, 0)
	cnt = make([]int, N)
	for i := 0; i < N; i++ {
		if len(node[i]) == 1 {
			q = append(q, i)
		}
		cnt[i] = len(node[i])
	}

	ans = make([]int, 0)
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		if cnt[cur] == 0 {
			continue
		}
		cnt[cur]--
		dfs(cur, -1, cur)
		// out("->", cur, cnt, p)

		for _, e := range node[cur] {
			if cnt[e.to] != 0 {
				cnt[e.to]--
			}
			if cnt[e.to] == 1 {
				q = append(q, e.to)
			}
		}
	}
	// out(q)
	out(len(ans))
	for _, e := range ans {
		fmt.Fprint(wr, e, " ")
	}
	out()
}
