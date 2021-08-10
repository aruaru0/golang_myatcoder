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
	a, x, y int
}

type priorityQueue []pqi

func (pq priorityQueue) Len() int      { return len(pq) }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].a < pq[j].a
}
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

var H, W int
var s []string

const inf = int(1e18)

func check(x, y int, crush [][]bool) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			px := x + i
			py := y + j
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if crush[py][px] == true {
				return true
			}
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
	H, W = getI(), getI()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
		}
	}
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, 0, 0})
	dist[0][0] = 0

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for len(pq) != 0 {
		c := pq[0]
		// out(c)
		heap.Pop(&pq)
		// if dist[c.y][c.x] < c.a {
		// 	continue
		// }
		for i := 0; i < 4; i++ {
			px := c.x + dx[i]
			py := c.y + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if s[py][px] == '#' {
				for y := -1; y <= 1; y++ {
					for x := -1; x <= 1; x++ {
						if x == 0 && y == 0 {
							continue
						}
						ppx := px + x
						ppy := py + y
						if ppx < 0 || ppx >= W || ppy < 0 || ppy >= H {
							continue
						}
						if dist[ppy][ppx] > dist[c.y][c.x]+1 {
							dist[ppy][ppx] = dist[c.y][c.x] + 1
							heap.Push(&pq, pqi{dist[ppy][ppx], ppx, ppy})
						}
					}
				}
			} else {
				if dist[py][px] > dist[c.y][c.x] {
					dist[py][px] = dist[c.y][c.x]
					heap.Push(&pq, pqi{dist[py][px], px, py})
				}
			}
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }

	out(dist[H-1][W-1])
}
