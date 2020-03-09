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

// Priority Queue
type Item struct {
	priority, value, index int
}

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

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// End Priority Queue
// Path
type Edge struct {
	to, cost int
}

type Path struct {
	edges []Edge
}

// Dijkstra
type Route struct {
	path []int
}

func Dijkstra(N, S int, path []Path) ([]int, []Route) {
	pq := make(PQ, 0)
	heap.Init(&pq)
	d := make([]int, N+1)
	r := make([]Route, N+1)
	// init
	for i := 0; i <= N; i++ {
		d[i] = math.MaxInt32
	}
	d[S] = 0
	r[S].path = []int{S}
	heap.Push(&pq, &Item{0, S, 0})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		v := item.value
		if d[v] < item.priority {
			continue
		}
		for _, e := range path[v].edges {
			if d[e.to] > d[v]+e.cost {
				d[e.to] = d[v] + e.cost
				r[e.to].path = append(r[v].path, e.to)
				heap.Push(&pq, &Item{d[e.to], e.to, 0})
			}
		}
	}
	return d, r

}

// テキスト "#"が壁のやつを幅優先探索(0,0)の距離
type queue struct {
	x, y int
}

func bfs(H, W int, s []string) [][]int {
	dist := make([][]int, H)
	// 初期化
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = -1
		}
	}
	q := make([]queue, 0)
	q = append(q, queue{0, 0})

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	dist[0][0] = 0
	for len(q) != 0 {
		x := q[0].x
		y := q[0].y
		q = q[1:]
		for k := 0; k < 4; k++ {
			yy := y + dy[k]
			xx := x + dx[k]
			if yy < 0 || xx < 0 || yy >= H || xx >= W {
				continue
			}
			if s[yy][xx] == '#' {
				continue
			}
			if dist[yy][xx] != -1 {
				continue
			}
			dist[yy][xx] = dist[y][x] + 1
			q = append(q, queue{xx, yy})
		}

	}

	return dist
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W := getInt(), getInt()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}

	ret := bfs(H, W, s)

	cnt := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] == '#' {
				cnt++
			}
		}
	}

	if ret[H-1][W-1] == -1 {
		out(-1)
	} else {
		out(H*W - ret[H-1][W-1] - cnt - 1)
	}

	/* ダイクストラ法での解法
	path := make([]Path, W*H)
	cnt := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] == '#' {
				cnt++
				continue
			}
			f := y*W + x
			if x != 0 && s[y][x-1] != '#' {
				t := f - 1
				path[f].edges = append(path[f].edges, Edge{t, 1})
			}
			if x != W-1 && s[y][x+1] != '#' {
				t := f + 1
				path[f].edges = append(path[f].edges, Edge{t, 1})
			}
			if y != 0 && s[y-1][x] != '#' {
				t := f - W
				path[f].edges = append(path[f].edges, Edge{t, 1})
			}
			if y != H-1 && s[y+1][x] != '#' {
				t := f + W
				path[f].edges = append(path[f].edges, Edge{t, 1})
			}
		}
	}
	d, _ := Dijkstra(W*H, 0, path)
	pos := W*H - 1


	if d[pos] == math.MaxInt32 {
		out(-1)
	} else {
		out(W*H - cnt - d[pos] - 1)
	}
	*/
}
