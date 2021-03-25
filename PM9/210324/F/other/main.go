package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

func scanString() string { sc.Scan(); return sc.Text() }
func scanRunes() []rune  { return []rune(scanString()) }
func scanInt() int       { a, _ := strconv.Atoi(scanString()); return a }
func scanInt64() int64   { a, _ := strconv.ParseInt(scanString(), 10, 64); return a }
func scanFloat() float64 { a, _ := strconv.ParseFloat(scanString(), 64); return a }

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func debug(a ...interface{}) {
	if os.Getenv("ONLINE_JUDGE") == "false" {
		fmt.Fprintln(os.Stderr, a...)
	}
}

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

//•*¨*•.¸¸♪main•*¨*•.¸¸♪(　-ω-)ノ　(　・ω・)
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1001), 1001001)

	n := scanInt()
	a := scanInts(n)
	b := scanInts(n)
	p := scanInts(n)
	for i := 0; i < n; i++ {
		p[i]--
	}

	ans := [][2]int{}
	pq := &priorityQueue{}
	for i := 0; i < n; i++ {
		*pq = append(*pq, S{i, p[i], b[p[i]]})
	}
	heap.Init(pq)

	for pq.Len() != 0 {
		nxt := heap.Pop(pq).(S)
		i, j := nxt.i, nxt.j
		if i == j || j != p[i] || a[i] <= b[j] || a[j] <= b[p[j]] {
			continue
		}
		ans = append(ans, [2]int{i, j})
		p[i], p[j] = p[j], j
		fmt.Println(p)
		heap.Push(pq, S{i, p[i], b[p[i]]})
	}

	for i := 0; i < n; i++ {
		if p[i] != i {
			fmt.Fprintln(wr, -1)
			return
		}
	}

	fmt.Fprintln(wr, len(ans))
	for _, v := range ans {
		fmt.Fprintln(wr, v[0]+1, v[1]+1)
	}

}

type S struct{ i, j, b int }

type priorityQueue []S

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].b > pq[j].b }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(S)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}
