package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

// Priority Queue
type Item struct {
	priority, pos, index int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].pos < pq[j].pos
	}
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

func main() {
	sc.Split(bufio.ScanWords)

	N, K, D := getInt(), getInt(), getInt()

	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	pq := make(PQ, 0)
	heap.Init(&pq)

	// ギリギリの末尾を計算
	n := (K - 1) * D
	pos := N - n - 1
	if pos < 0 {
		out(-1)
		return
	}
	idx := 0
	ans := make([]int, 0)
	cur := -1
	for {
		for idx < N {
			if idx <= pos {
				heap.Push(&pq, &Item{a[idx], idx, 0})
				idx++
			} else {
				break
			}
		}
		var item *Item
		for {
			item = heap.Pop(&pq).(*Item)
			if item.pos >= cur {
				break
			}
		}
		cur = item.pos + D //　手前を更新
		ans = append(ans, item.priority)
		pos += D // 末尾を更新
		K--
		if K == 0 {
			break
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, v := range ans {
		fmt.Fprint(w, v, " ")
	}
	fmt.Fprintln(w)
}
