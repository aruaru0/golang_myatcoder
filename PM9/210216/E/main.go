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

type pos struct {
	y, x int
}

const inf = int(1e18)

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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	a := make([]string, H)
	for i := 0; i < H; i++ {
		a[i] = getS()
	}
	warp := make(map[byte][]pos)
	var start, goal pos

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if 'a' <= a[i][j] && a[i][j] <= 'z' {
				warp[a[i][j]] = append(warp[a[i][j]], pos{i, j})
			}
			if a[i][j] == 'S' {
				start = pos{i, j}
			}
			if a[i][j] == 'G' {
				goal = pos{i, j}
			}
		}
	}

	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
		}
	}
	used := make([]bool, 26)
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	dist[start.y][start.x] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, start.x, start.y})
	for len(pq) != 0 {
		c := pq[0]
		heap.Pop(&pq)
		if dist[c.y][c.x] < c.a {
			continue
		}
		for i := 0; i < 4; i++ {
			px := c.x + dx[i]
			py := c.y + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if a[py][px] == '#' {
				continue
			}
			if 'a' <= a[py][px] && a[py][px] <= 'z' {
				x := int(a[py][px] - 'a')
				if used[x] != true {
					next, ok := warp[a[py][px]]
					if ok {
						for _, e := range next {
							if e.x == px && e.y == py {
								continue
							}
							if dist[e.y][e.x] > dist[c.y][c.x]+2 {
								dist[e.y][e.x] = dist[c.y][c.x] + 2
								heap.Push(&pq, pqi{dist[e.y][e.x], e.x, e.y})
							}
						}
					}
					used[x] = true
				}
			}

			if dist[py][px] > dist[c.y][c.x]+1 {
				dist[py][px] = dist[c.y][c.x] + 1
				heap.Push(&pq, pqi{dist[py][px], px, py})
			}
		}
	}
	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }
	ans := dist[goal.y][goal.x]
	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
