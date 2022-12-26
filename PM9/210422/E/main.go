package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string   { sc.Scan(); return sc.Text() }
func scanRunes() []rune    { return []rune(scanString()) }
func scanInt() int         { a, _ := strconv.Atoi(scanString()); return a }
func scanInt64() int64     { a, _ := strconv.ParseInt(scanString(), 10, 64); return a }
func scanFloat64() float64 { a, _ := strconv.ParseFloat(scanString(), 64); return a }

func scanInts(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = scanInt()
	}
	return res
}

func debug(a ...interface{}) { fmt.Fprintln(os.Stderr, a...) }

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func out(x ...interface{}) {
	fmt.Println(x...)
}

type pqi struct{ a, i int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].a > pq[j].a }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

type pqii struct{ a, i, d int }

type priorityQueue2 []pqii

func (pq priorityQueue2) Len() int            { return len(pq) }
func (pq priorityQueue2) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue2) Less(i, j int) bool  { return pq[i].a < pq[j].a }
func (pq *priorityQueue2) Push(x interface{}) { *pq = append(*pq, x.(pqii)) }
func (pq *priorityQueue2) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

// priority queueの実装を知りたいのでコードをコピー
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	all := 200000
	n, q := scanInt(), scanInt()
	a := make([]int, n)
	b := make([]int, n)
	pqplace := make([]priorityQueue, all)
	now := make([]int, n)

	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt(), scanInt()-1
		heap.Push(&pqplace[b[i]], pqi{a[i], i})
		now[i] = b[i]
	}

	pq := priorityQueue2{}
	for i := 0; i < all; i++ {
		if len(pqplace[i]) == 0 {
			continue
		}
		t := pqplace[i][0]
		heap.Push(&pq, pqii{t.a, t.i, i})
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < q; i++ {
		c, d := scanInt()-1, scanInt()-1

		src := now[c]
		now[c] = d

		// 先頭から場所が異なるものを消して、一番上を登録
		// 移動前の場所
		for len(pqplace[src]) != 0 {
			item := pqplace[src][0]
			if now[item.i] == src {
				heap.Push(&pq, pqii{item.a, item.i, src})
				break
			}
			heap.Pop(&pqplace[src])
		}

		heap.Push(&pqplace[d], pqi{a[c], c})

		// 先頭から場所が異なるものを消して、一番上を登録
		// 移動先の場所
		for len(pqplace[d]) != 0 {
			item := pqplace[d][0]
			if now[item.i] == d {
				heap.Push(&pq, pqii{item.a, item.i, d})
				break
			}
			heap.Pop(&pqplace[d])
		}

		for {
			//　現在の場所と一致していないものを削除して、一番上を表示
			mi := pq[0]
			if len(pqplace[mi.d]) != 0 && pqplace[mi.d][0].a == mi.a {
				fmt.Fprintln(wr, mi.a)
				break
			}
			heap.Pop(&pq)
		}
	}
}
