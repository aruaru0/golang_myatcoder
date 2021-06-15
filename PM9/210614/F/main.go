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

type pqi struct{ a, x, y int }

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

type pos struct {
	x, y int
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, K := getI(), getI(), getI()
	y1, x1 := getI()-1, getI()-1
	y2, x2 := getI()-1, getI()-1
	c := make([]string, H)
	for i := 0; i < H; i++ {
		c[i] = getS()
	}

	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, x1, y1})
	dist[y1][x1] = 0
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if dist[cur.y][cur.x] < cur.a {
			continue
		}
		for i := 0; i < 4; i++ {
			px := cur.x
			py := cur.y
			for k := 0; k < K; k++ {
				px += dx[i]
				py += dy[i]
				if px < 0 || px >= W || py < 0 || py >= H {
					break
				}
				if c[py][px] == '@' {
					break
				}
				if dist[py][px] < dist[cur.y][cur.x]+1 {
					break
				}
				if dist[py][px] > dist[cur.y][cur.x]+1 {
					dist[py][px] = dist[cur.y][cur.x] + 1
					heap.Push(&pq, pqi{dist[py][px], px, py})
				}
			}
		}
	}

	ans := dist[y2][x2]
	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
