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

type pos struct {
	i, j, dir int
}

type dirc struct {
	d0, d1 int
	c      byte
}

var dirs = map[dirc]bool{
	//	 1
	// 0   2
	//	 3
	{0, 2, 'A'}: true,
	{1, 3, 'A'}: true,
	{2, 0, 'A'}: true,
	{3, 1, 'A'}: true,

	{0, 3, 'B'}: true,
	{3, 0, 'B'}: true,
	{1, 2, 'B'}: true,
	{2, 1, 'B'}: true,

	{0, 1, 'C'}: true,
	{1, 0, 'C'}: true,
	{2, 3, 'C'}: true,
	{3, 2, 'C'}: true,
}

func calc(dir0, dir1 int, c byte) int {
	if dirs[dirc{dir0, dir1, c}] {
		return 0
	}
	return 1
}

type pqi struct{ cost, i, j, d int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

func solve() {
	H, W := getI(), getI()
	s := getStrings(H)
	const inf = int(1e18)
	// dp[i][j][k] = (i, j)に方向kから到達するまでの最小のコスト（切り替え数）
	dp := make([][][4]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([][4]int, W)
		for j := 0; j < W; j++ {
			for k := 0; k < 4; k++ {
				dp[i][j][k] = inf
			}
		}
	}

	//   1
	// 0   2
	//   3
	dj := []int{1, 0, -1, 0}
	di := []int{0, 1, 0, -1}
	curidx := []int{2, 3, 0, 1}

	dp[0][0][0] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, 0, 0, 0})

	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if cur.cost > dp[cur.i][cur.j][cur.d] {
			continue
		}
		for i := 0; i < 4; i++ {
			ni, nj := cur.i+di[i], cur.j+dj[i]
			if ni < 0 || ni >= H || nj < 0 || nj >= W {
				continue
			}
			// out(cur, i, "next= ", ni, nj, string(s[cur.i][cur.j]))
			cost := calc(cur.d, curidx[i], s[cur.i][cur.j])
			if dp[ni][nj][i] > dp[cur.i][cur.j][cur.d]+cost {
				dp[ni][nj][i] = dp[cur.i][cur.j][cur.d] + cost
				heap.Push(&pq, pqi{dp[ni][nj][i], ni, nj, i})
			}
		}
	}

	// out(dp[H-1][W-1])

	ans := min(dp[H-1][W-1][0]+calc(0, 2, s[H-1][W-1]),
		dp[H-1][W-1][1]+calc(1, 2, s[H-1][W-1]))
	out(ans)

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
		// break
	}
}
