package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
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

func getString() string {
	sc.Scan()
	return sc.Text()
}

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

// priority Queue----------------------------

// Item :
type Item struct {
	priority, x, y, index int
}

// PQ :
type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push :
func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop :
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// End Priority Queue ----------------------

//
// ２次元のマス目のダイクストラ法
//  sx,sy: 開始位置 H,W:幅と高さ e: [H][W]のコスト
//  戻り値: sx,syからの最小コスト
//
func dijkstra(sx, sy, H, W int, e [][]int) [][]int {
	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = math.MaxInt32
		}
	}

	pq := make(PQ, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{0, sx, sy, 0})
	dist[sy][sx] = 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		x, y := item.x, item.y
		if dist[y][x] < item.priority {
			continue
		}
		dx := []int{0, -1, 1, 0}
		dy := []int{-1, 0, 0, 1}
		for i := 0; i < 4; i++ {
			xx := x + dx[i]
			yy := y + dy[i]
			if xx < 0 || yy < 0 || xx >= W || yy >= H {
				continue
			}
			if dist[yy][xx] > dist[y][x]+e[yy][xx] {
				dist[yy][xx] = dist[y][x] + e[yy][xx]
				heap.Push(&pq, &Item{dist[yy][xx], xx, yy, 0})
			}
		}

	}

	return dist
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W := getInt(), getInt()
	e := make([][]int, H)
	for i := 0; i < H; i++ {
		e[i] = make([]int, W)
		for j := 0; j < W; j++ {
			e[i][j] = getInt()
		}
	}

	for i := 0; i < H; i++ {
		out(e[i])
	}

	ret := dijkstra(0, 0, H, W, e)
	out("----------------------")
	for i := 0; i < H; i++ {
		out(ret[i])
	}

}
