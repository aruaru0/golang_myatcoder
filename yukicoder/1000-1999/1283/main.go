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

type pqi struct{ a, x, y, t int }

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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
	}
	for i := 0; i < M; i++ {
		h, w, c := getI()-1, getI()-1, getI()
		a[h][w] = c
	}

	d := make([][][2]int, N)
	for i := 0; i < N; i++ {
		d[i] = make([][2]int, N)
		for j := 0; j < N; j++ {
			d[i][j][0] = inf
			d[i][j][1] = inf
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, 0, 0, 0})
	d[0][0][0] = 0
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		for i := 0; i < 4; i++ {
			px := cur.x + dx[i]
			py := cur.y + dy[i]
			if px < 0 || px >= N || py < 0 || py >= N {
				continue
			}
			if cur.t == 0 && a[py][px] != 0 {
				if d[py][px][1] > d[cur.y][cur.x][cur.t]+1 {
					d[py][px][1] = d[cur.y][cur.x][cur.t] + 1
					heap.Push(&pq, pqi{d[py][px][1], px, py, 1})
				}
			}
			if d[py][px][cur.t] > d[cur.y][cur.x][cur.t]+a[py][px]+1 {
				d[py][px][cur.t] = d[cur.y][cur.x][cur.t] + a[py][px] + 1
				heap.Push(&pq, pqi{d[py][px][1], px, py, cur.t})
			}
		}
	}
	// for i := 0; i < N; i++ {
	// 	out(a[i])
	// }
	// out("-----")
	// for i := 0; i < N; i++ {
	// 	out(d[i])
	// }
	out(min(d[N-1][N-1][0], d[N-1][N-1][1]))
}
