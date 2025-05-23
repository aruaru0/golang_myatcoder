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

type BIT struct {
	v []int
}

func newBIT(n int) *BIT {
	return &BIT{
		v: make([]int, n+2), // 1-based indexing
	}
}

func (b *BIT) sum(a int) int {
	res := 0
	for i := a; i > 0; i -= i & -i {
		res += b.v[i]
	}
	return res
}

func (b *BIT) rangeSum(x, y int) int {
	if x > y {
		return 0
	}
	return b.sum(y) - b.sum(x-1)
}

func (b *BIT) add(a, w int) {
	for i := a; i < len(b.v); i += i & -i {
		b.v[i] += w
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	N := getI()
	edgesInput := make([][2]int, N-1)
	adj := make([][]int, N+1)

	for i := 0; i < N-1; i++ {
		u, v := getI(), getI()
		edgesInput[i] = [2]int{u, v}
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	parent := make([]int, N+1)
	children := make([][]int, N+1)

	// BFS to build parent and children
	q := []int{1}
	parent[1] = 0
	visited := make([]bool, N+1)
	visited[1] = true

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adj[u] {
			if !visited[v] && v != parent[u] {
				parent[v] = u
				children[u] = append(children[u], v)
				visited[v] = true
				q = append(q, v)
			}
		}
	}

	out(parent, children)

	inTime := make([]int, N+1)
	outTime := make([]int, N+1)
	timeCounter := 1

	// Recursive DFS to compute in and out time
	var dfs func(u int)
	dfs = func(u int) {
		inTime[u] = timeCounter
		timeCounter++
		for _, v := range children[u] {
			dfs(v)
		}
		outTime[u] = timeCounter - 1
	}

	dfs(1)

	out(inTime, outTime)
	// For each edge, determine which is the child node
	edgeChild := make([]int, N) // 1-based index for edges
	for i := 0; i < N-1; i++ {
		u, v := edgesInput[i][0], edgesInput[i][1]
		if parent[u] == v {
			edgeChild[i+1] = u
		} else if parent[v] == u {
			edgeChild[i+1] = v
		}
	}

	out(edgeChild)
	// Initialize BIT with all 1s
	BITSize := N + 2
	bit := newBIT(BITSize)
	for i := 1; i <= N; i++ {
		bit.add(inTime[i], 1)
	}
	totalSum := N

	Q := getI()
	output := make([]string, 0)

	for q := 0; q < Q; q++ {
		t := getI()
		if t == 1 {
			x, w := getI(), getI()
			bit.add(inTime[x], w)
			totalSum += w
		} else {
			y := getI()
			c := edgeChild[y]
			l, r := inTime[c], outTime[c]
			sumC := bit.rangeSum(l, r)
			diff := abs(2*sumC - totalSum)
			output = append(output, fmt.Sprintf("%d", diff))
		}
	}

	for _, s := range output {
		fmt.Println(s)
	}
}
