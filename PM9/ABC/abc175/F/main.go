package main

import (
	"bufio"
	"container/heap"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	s := make([]string, n)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		s[i], c[i] = getS(), getI()
	}

	reverse := func(s string) string {
		res := ""
		for i := 0; i < len(s); i++ {
			res += string(s[len(s)-1-i])
		}
		return res
	}

	for i := 0; i < n; i++ {
		s = append(s, reverse(s[i]))
		c = append(c, c[i])
	}

	id := make([][]int, n*2)
	m := 1
	for i := 0; i < n*2; i++ {
		id[i] = make([]int, len(s[i]))
		for j := 0; j < len(s[i]); j++ {
			id[i][j] = m
			m++
		}
	}

	isPal := func(s string) bool {
		f := true
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				f = false
			}
		}
		return f
	}

	g := make([][]edge, m)
	for i := 0; i < n*2; i++ {
		for j := 0; j < len(s[i]); j++ {
			for k := 0; k < n*2; k++ {
				if i/n == k/n {
					continue
				}
				w := min(len(s[k]), len(s[i][j:]))
				if s[i][j:j+w] != s[k][:w] {
					continue
				}

				u := 0
				if len(s[k]) > w {
					u = id[k][w]
					if isPal(s[k][w:]) {
						u = 0
					}
				}
				if len(s[i])-j > w {
					u = id[i][j+w]
					if isPal(s[i][j+w:]) {
						u = 0
					}
				}
				v := id[i][j]
				g[v] = append(g[v], edge{u, c[k]})
			}
		}
	}

	ans := inf
	for i := 0; i < n; i++ {
		co := dijkstra(g, id[i][0])[0]
		if co == inf {
			continue
		}
		ans = min(ans, co+c[i])
	}
	for i := 0; i < n; i++ {
		if isPal(s[i]) {
			ans = min(ans, c[i])
		}
	}

	if ans == inf {
		fmt.Fprintln(wr, -1)
	} else {
		fmt.Fprintln(wr, ans)
	}

}

const inf = 1<<63 - 1

type edge struct{ to, co int }
type vert struct{ d, v int }
type pqForDijkstra []vert

func (pq pqForDijkstra) Len() int            { return len(pq) }
func (pq pqForDijkstra) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq pqForDijkstra) Less(i, j int) bool  { return pq[i].d < pq[j].d }
func (pq *pqForDijkstra) Push(x interface{}) { *pq = append(*pq, x.(vert)) }
func (pq *pqForDijkstra) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

func dijkstra(es [][]edge, s int) []int {
	d := make([]int, len(es))
	for i := 0; i < len(es); i++ {
		d[i] = inf
	}
	d[s] = 0
	pq := &pqForDijkstra{vert{0, s}}
	heap.Init(pq)
	for len(*pq) != 0 {
		p := heap.Pop(pq).(vert)
		if p.d > d[p.v] {
			continue
		}
		for _, e := range es[p.v] {
			cost := d[p.v] + e.co
			if cost < d[e.to] {
				d[e.to] = cost
				heap.Push(pq, vert{d[e.to], e.to})
			}
		}
	}
	return d
}
