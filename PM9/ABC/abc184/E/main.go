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

func pos2idx(x, y int) int {
	return x + y*W
}

func idx2pos(idx int) (int, int) {
	return idx % W, idx / W
}

var H, W int

const inf = int(1e18)

type pqi struct{ a, to int }

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
	H, W = getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	node := make([][]int, H*W+26)
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	start, goal := 0, 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if s[h][w] == 'S' {
				start = pos2idx(w, h)
			} else if s[h][w] == 'G' {
				goal = pos2idx(w, h)
			} else if s[h][w] >= 'a' && s[h][w] <= 'z' {
				node[pos2idx(w, h)] = append(node[pos2idx(w, h)], W*H+int(s[h][w]-'a'))
				node[W*H+int(s[h][w]-'a')] = append(node[W*H+int(s[h][w]-'a')], pos2idx(w, h))
			}
			for i := 0; i < 4; i++ {
				px := w + dx[i]
				py := h + dy[i]
				if px < 0 || px >= W || py < 0 || py >= H {
					continue
				}
				if s[py][px] == '#' {
					continue
				}
				node[pos2idx(w, h)] = append(node[pos2idx(w, h)], pos2idx(px, py))
			}
		}
	}

	dist := make([]int, H*W+26)
	for i := 0; i < H*W+26; i++ {
		dist[i] = inf
	}
	dist[start] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, start})
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if dist[cur.to] < cur.a {
			continue
		}
		for _, nxt := range node[cur.to] {
			d := 2
			if nxt >= H*W || cur.to >= H*W {
				d = 1
			}
			if dist[nxt] > dist[cur.to]+d {
				dist[nxt] = dist[cur.to] + d
				heap.Push(&pq, pqi{dist[nxt], nxt})
			}
		}
	}
	if dist[goal] == inf {
		fmt.Println(-1)
		return
	} else {
		out(dist[goal] / 2)
	}
}
