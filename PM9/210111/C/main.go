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

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H)
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
		a[i] = make([]int, W)
		for j := 0; j < W; j++ {
			a[i][j] = inf
		}
	}

	dx := []int{1, 0}
	dy := []int{0, 1}

	pq := priorityQueue{}
	if s[0][0] == '.' {
		heap.Push(&pq, pqi{0, 0, 0})
		a[0][0] = 0
	} else {
		heap.Push(&pq, pqi{1, 0, 0})
		a[0][0] = 1
	}
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		for i := 0; i < 2; i++ {
			px := cur.x + dx[i]
			py := cur.y + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			cost := 0
			if s[cur.y][cur.x] == '.' && s[py][px] == '#' {
				cost++
			}
			if a[py][px] > a[cur.y][cur.x]+cost {
				a[py][px] = a[cur.y][cur.x] + cost
				heap.Push(&pq, pqi{a[py][px], px, py})
			}
		}
	}

	ans := a[H-1][W-1]
	out(ans)
}
