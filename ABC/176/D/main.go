package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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
var ch, cw int
var dh, dw int
var S []string
var u [][]int

//---------------------------------------------
// priority queue
//---------------------------------------------
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

const inf = int(1e9)

func bsf(sx, sy int) int {
	u[sy][sx] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, sx, sy})
	for len(pq) != 0 {
		cx := pq[0].x
		cy := pq[0].y
		heap.Pop(&pq)
		for dy := -2; dy <= 2; dy++ {
			for dx := -2; dx <= 2; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				x := cx + dx
				y := cy + dy
				if x < 0 || x >= W || y < 0 || y >= H {
					continue
				}
				if S[y][x] == '#' {
					continue
				}
				p := u[cy][cx]
				if (dx == 0 && abs(dy) == 1) || (dy == 0 && abs(dx) == 1) {
				} else {
					p++
				}
				if u[y][x] <= p {
					continue
				}
				u[y][x] = p
				heap.Push(&pq, pqi{p, x, y})
			}
		}
	}
	return 1
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	H, W = getInt(), getInt()
	ch, cw = getInt()-1, getInt()-1
	dh, dw = getInt()-1, getInt()-1
	S = make([]string, H)
	u = make([][]int, H)
	for i := 0; i < H; i++ {
		S[i] = getString()
		u[i] = make([]int, W)
		for j := 0; j < W; j++ {
			u[i][j] = inf
		}
	}

	bsf(cw, ch)

	// for i := 0; i < H; i++ {
	// 	out(u[i])
	// }

	if u[dh][dw] == inf {
		out(-1)
		return
	}
	out(u[dh][dw])
}
