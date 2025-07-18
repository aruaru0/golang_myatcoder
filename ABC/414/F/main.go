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

const INF = int(1e18)

type Edge struct {
	U, V int
}

type Pair struct {
	To     int
	EdgeID int
}

type QueueItem struct {
	EdgeID int
	Dist   int
}

func solve() {
	n := getI()
	k := getI()

	edges := make([]Edge, 0, n-1)
	g := make([][]Pair, n)

	for i := 0; i < n-1; i++ {
		a := getI() - 1
		b := getI() - 1

		edges = append(edges, Edge{a, b})

		g[a] = append(g[a], Pair{b, i})
		g[b] = append(g[b], Pair{a, n - 1 + i})
	}

	dist := make([][]int, (n-1)*2)
	for i := range dist {
		dist[i] = make([]int, k)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}

	cnt := make([][]int, n)
	for i := range cnt {
		cnt[i] = make([]int, k)
	}

	q := []QueueItem{}

	push := func(ei, d int) {
		w := d % k
		if dist[ei][w] != INF {
			return
		}
		dist[ei][w] = d
		q = append(q, QueueItem{ei, d})
	}

	for _, entry := range g[0] {
		push(entry.EdgeID, 1)
	}

	for len(q) != 0 {
		item := q[0]
		q = q[1:]

		ei := item.EdgeID
		d := item.Dist

		edgePair := edges[ei%(n-1)]
		var v int
		if ei < n-1 {
			v = edgePair.V
		} else {
			v = edgePair.U
		}

		cnt[v][d%k]++

		if cnt[v][d%k] <= 2 {
			for _, entry := range g[v] {
				ej := entry.EdgeID
				if d%k != 0 && ei%(n-1) == ej%(n-1) {
					continue
				}
				push(ej, d+1)
			}
		}
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = INF
	}

	for i := 0; i < (n-1)*2; i++ {
		edgePair := edges[i%(n-1)]
		a, b := edgePair.U, edgePair.V
		if i >= n-1 {
			a, b = b, a
		}
		ans[b] = min(ans[b], dist[i][0])
	}

	resultSlice := make([]int, 0, n-1)
	for i := 1; i < n; i++ {
		if ans[i] == INF {
			resultSlice = append(resultSlice, -1)
		} else {
			resultSlice = append(resultSlice, ans[i]/k)
		}
	}
	outSlice(resultSlice)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()

	for i := 0; i < T; i++ {
		solve()
	}

}
