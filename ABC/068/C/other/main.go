package main

import (
	"bufio"
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

// Queue ... 実装例
type QueueL[T any] struct {
	data []T
	less func(i, j T) bool
}

func New[T any](less func(i, j T) bool) QueueL[T] {
	var ret QueueL[T]
	ret.data = make([]T, 0)
	ret.less = less
	return ret
}

func (q QueueL[T]) Len() int {
	return len(q.data)
}

func (q QueueL[T]) Swap(i, j int) {
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

func (q *QueueL[T]) Push(x T) {
	q.data = append(q.data, x)
	cur := q.Len()
	parent := cur / 2
	for cur != 1 {
		if q.less(q.data[cur-1], q.data[parent-1]) {
			q.Swap(cur-1, parent-1)
		} else {
			break
		}
		cur = parent
		parent = cur / 2
	}
}

func (q *QueueL[T]) Pop() (T, bool) {
	if q.Len() == 0 {
		var item T
		return item, false
	}
	old := *q
	n := len(old.data)
	item := old.data[0]

	old.data[0] = old.data[n-1]
	old.data = old.data[:n-1]
	cur := 1
	for {
		nxt0 := cur * 2
		nxt1 := cur*2 + 1
		if nxt0 > len(old.data) {
			break
		}
		nxt := nxt0
		if nxt1 <= len(old.data) && old.less(q.data[nxt1-1], q.data[nxt0-1]) {
			nxt = nxt1
		}
		if old.less(q.data[nxt-1], q.data[cur-1]) {
			old.Swap(nxt-1, cur-1)
		} else {
			break
		}

		cur = nxt
	}

	*q = old
	return item, true
}

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
	pq := New[Item](func(i, j Item) bool {
		return i.priority < j.priority
	})
	// heap.Init(&pq)
	d := make([]int, N+1)
	r := make([]Route, N+1)
	// init
	for i := 0; i <= N; i++ {
		d[i] = math.MaxInt32
	}
	d[S] = 0
	r[S].path = []int{S}
	pq.Push(Item{0, S, 0})
	for pq.Len() > 0 {
		item, _ := pq.Pop()
		v := item.value
		if d[v] < item.priority {
			continue
		}
		for _, e := range path[v].edges {
			if d[e.to] > d[v]+e.cost {
				d[e.to] = d[v] + e.cost
				r[e.to].path = append(r[v].path, e.to)
				pq.Push(Item{d[e.to], e.to, 0})
			}
		}
	}
	return d, r

}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	path := make([]Path, N+1)

	for i := 0; i < M; i++ {
		f, t := getInt(), getInt()
		path[f].edges = append(path[f].edges, Edge{t, 1})
		path[t].edges = append(path[t].edges, Edge{f, 1})
	}

	S, D := 1, N

	d, _ := Dijkstra(N, S, path)

	if d[D] == 2 {
		out("POSSIBLE")
	} else {
		out("IMPOSSIBLE")
	}
}
