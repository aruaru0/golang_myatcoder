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

type pqi struct{ a, y, x int }

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
	H, W, Y := getI(), getI(), getI()
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = getInts(W)
	}

	pq := priorityQueue{}
	used := make([][]bool, H)
	for i := 0; i < H; i++ {
		used[i] = make([]bool, W)
	}

	for i := 0; i < H; i++ {
		heap.Push(&pq, pqi{a[i][0], i, 0})
		heap.Push(&pq, pqi{a[i][W-1], i, W - 1})
		used[i][0] = true
		used[i][W-1] = true
	}
	for i := 1; i < W-1; i++ {
		heap.Push(&pq, pqi{a[0][i], 0, i})
		heap.Push(&pq, pqi{a[H-1][i], H - 1, i})
		used[0][i] = true
		used[H-1][i] = true
	}

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	tot := H * W

	check := make([][]int, H)
	for i := 0; i < H; i++ {
		check[i] = make([]int, W)
	}

	// return
	for y := 1; y <= Y; y++ {
		// out("y =", y)
		for len(pq) != 0 && pq[0].a <= y {
			cur := pq[0]
			heap.Pop(&pq)
			if check[cur.y][cur.x] == 0 {
				tot--
			}
			check[cur.y][cur.x]++
			// out(cur)
			for i := 0; i < 4; i++ {
				px := cur.x + dx[i]
				py := cur.y + dy[i]
				if py < 0 || py >= H || px < 0 || px >= W {
					continue
				}
				if used[py][px] {
					continue
				}
				heap.Push(&pq, pqi{a[py][px], py, px})
				used[py][px] = true
			}
		}
		// for i := 0; i < H; i++ {
		// 	out(check[i])
		// }
		out(tot)
	}

}
