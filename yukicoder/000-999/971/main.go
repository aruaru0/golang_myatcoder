package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

var H, W int
var s []string

var dist [][]int

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ d, cnt, x, y int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].d < pq[j].d }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

var dx = []int{1, 0}
var dy = []int{0, 1}

func dijkstra(sx, sy int) {
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, 0, sx, sy})
	dist[sx][sy] = 0
	for len(pq) != 0 {
		cnt, cx, cy := pq[0].cnt, pq[0].x, pq[0].y
		heap.Pop(&pq)
		cnt++
		for i := 0; i < 2; i++ {
			px := cx + dx[i]
			py := cy + dy[i]
			if px >= W || py >= H {
				continue
			}
			d := 1
			if s[py][px] == 'k' {
				d += cnt
			}
			// out("px, py", px, py, string(s[py][px]), d)
			if dist[py][px] > dist[cy][cx]+d {
				dist[py][px] = dist[cy][cx] + d
				heap.Push(&pq, pqi{dist[py][px], cnt, px, py})
			}

		}
	}
}

const inf = int(1e15)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	dist = make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
		}
	}

	dijkstra(0, 0)
	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }
	out(dist[H-1][W-1])
}
