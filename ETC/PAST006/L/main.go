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

type pqi struct {
	a    float64
	x, y int
}

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].a < pq[j].a }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

func d(x0, y0, x1, y1 float64) float64 {
	dx := x0 - x1
	dy := y0 - y1
	return math.Hypot(dx, dy)
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

func f(M int, cx, cy, r []float64) float64 {

	pq := priorityQueue{}

	// tower to tower
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			heap.Push(&pq, pqi{d(px[i], py[i], px[j], py[j]), i, j})
		}
	}

	// loop to loop
	for i := 0; i < M; i++ {
		for j := i + 1; j < M; j++ {
			dist := d(cx[i], cy[i], cx[j], cy[j])
			r0, r1 := r[i], r[j]
			if r0 > r1 {
				r0, r1 = r1, r0
			}
			// out(i, j, "C", cx[i], cy[i], r[i], "d", cx[j], cy[j], r[j], "dist", dist, "r", r1, r0)
			if dist < r1-r0 { // 内包している
				dist = r1 - r0 - dist
				heap.Push(&pq, pqi{dist, i + N, j + N})
			} else if r1-r0 <= dist && dist <= r1+r0 { // ２点で交わる
				heap.Push(&pq, pqi{0, i + N, j + N})
			} else { // 外にある
				dist -= r1 + r0
				heap.Push(&pq, pqi{dist, i + N, j + N})
			}
		}
	}
	// tower to loop
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			dist := d(px[i], py[i], cx[j], cy[j])
			r0 := r[j]
			if dist < r0 {
				dist = r0 - dist
				heap.Push(&pq, pqi{dist, i, j + N})
			} else {
				dist -= r0
				heap.Push(&pq, pqi{dist, i, j + N})
			}
		}
	}

	uf := newDsu(N + M)
	ans := 0.0
	m := make(map[int]bool)
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if uf.Same(cur.x, cur.y) {
			continue
		} else {
			// out(cur)
			uf.Merge(cur.x, cur.y)
			ans += cur.a
			if cur.x < N && uf.Same(0, cur.x) {
				m[cur.x] = true
			}
			if cur.y < N && uf.Same(0, cur.y) {
				m[cur.y] = true
			}
		}
		if len(m) == N {
			break
		}
	}
	// out(m, cx, cy, r, ans)
	return ans
}

var N, M int
var px, py []float64

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	px = make([]float64, N)
	py = make([]float64, N)
	for i := 0; i < N; i++ {
		px[i], py[i] = getF(), getF()
	}
	cx := make([]float64, M)
	cy := make([]float64, M)
	r := make([]float64, M)
	for i := 0; i < M; i++ {
		cx[i], cy[i], r[i] = getF(), getF(), getF()
	}

	m := 1 << M
	ans := math.MaxFloat64
	for i := 0; i < m; i++ {
		x := make([]float64, 0)
		y := make([]float64, 0)
		rr := make([]float64, 0)
		for j := 0; j < M; j++ {
			if (i>>j)%2 == 1 {
				x = append(x, cx[j])
				y = append(y, cy[j])
				rr = append(rr, r[j])
			}
		}
		ans = math.Min(ans, f(len(x), x, y, rr))
	}

	out(ans)
}
