package main

import (
	"bufio"
	"container/list"
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

type pqi struct{ a, r, c int }

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
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}
	A, B, C, D := getI()-1, getI()-1, getI()-1, getI()-1

	const inf = int(1e9)
	p := make([][]int, H)
	used := make([][]bool, H)
	for i := 0; i < H; i++ {
		p[i] = make([]int, W)
		used[i] = make([]bool, W)
		for j := 0; j < W; j++ {
			p[i][j] = inf
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	q := list.New()

	push := func(r, c, d, cost int) {
		if p[r][c] <= d {
			return
		}
		p[r][c] = d
		if cost == 0 {
			q.PushFront(pqi{d, r, c})
		} else {
			q.PushBack(pqi{d, r, c})
		}
	}

	push(A, B, 0, 0)
	for q.Len() != 0 {
		r, c := q.Front().Value.(pqi).r, q.Front().Value.(pqi).c
		// out(r, c)
		q.Remove(q.Front())
		used[r][c] = true
		for i := 0; i < 4; i++ {
			nr, nc := r+dy[i], c+dx[i]
			if nr < 0 || nr >= H || nc < 0 || nc >= W {
				continue
			}
			if s[nr][nc] == '.' {
				push(nr, nc, p[r][c], 0)
			}
		}
		for i := 0; i < 4; i++ {
			nr, nc := r, c
			if nr < 0 || nr >= H || nc < 0 || nc >= W {
				continue
			}
			for j := 0; j < 2; j++ {
				nr += dy[i]
				nc += dx[i]
				if nr < 0 || nr >= H || nc < 0 || nc >= W {
					break
				}
				push(nr, nc, p[r][c]+1, 1)
			}
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(p[i])
	// }
	out(p[C][D])
}
