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

type pqi struct{ a, i int }

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

func withRotate() []int {
	p := make([]int, len(op))
	copy(p, op)
	if K == 0 {
		return op
	}

	// 一番小さいものが先頭に来るようにローテートする
	minpos := N - 1
	for i := N - K; i < N; i++ {
		if p[i] < p[minpos] {
			minpos = i
		}
	}
	r := N - minpos
	p = append(p[minpos:], p[:minpos]...)

	// 削除の処理
	ans := make([]int, 0)
	pq := priorityQueue{}

	for i := 0; i < K+1; i++ {
		heap.Push(&pq, pqi{p[i], i})
	}

	must := K + 1
	lastpos := -1
	for len(pq) != 0 {
		val, pos := pq[0].a, pq[0].i
		heap.Pop(&pq)
		if pos < lastpos {
			continue
		}
		ans = append(ans, val)
		lastpos = pos
		if lastpos >= r {
			if must < N {
				heap.Push(&pq, pqi{p[must], must})
				must++
			} else {
				break
			}
		}
	}
	return ans
}

func withoutRotate() []int {
	p := make([]int, len(op))
	copy(p, op)
	r := 0
	ans := make([]int, 0)
	pq := priorityQueue{}

	// K個まで優先キューに入れる
	for i := 0; i < K+1; i++ {
		heap.Push(&pq, pqi{p[i], i})

	}
	must := K + 1
	lastpos := -1 // 最後に取ったポイント
	for len(pq) != 0 {
		val, pos := pq[0].a, pq[0].i
		heap.Pop(&pq)
		if pos < lastpos {
			continue
		}
		ans = append(ans, val)
		lastpos = pos
		if lastpos >= r {
			if must < N { // 使ったら削除できるやつが増えるので補充
				heap.Push(&pq, pqi{p[must], must})
				must++
			} else {
				break
			}
		}
	}
	return ans
}

var N, K int
var op []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K = getI(), getI()
	op = getInts(N)

	x := withRotate()
	y := withoutRotate()
	n := min(len(x), len(y))
	sel := 0
	for i := 0; i < n; i++ {
		if x[i] < y[i] {
			sel = 1
			break
		}
		if x[i] > y[i] {
			sel = 2
			break
		}
	}
	if sel == 0 {
		if len(x) < len(y) {
			sel = 1
		} else {
			sel = 2
		}
	}

	if sel == 1 {
		for i := 0; i < len(x); i++ {
			fmt.Fprint(wr, x[i], " ")
		}
		out()
	} else {
		for i := 0; i < len(y); i++ {
			fmt.Fprint(wr, y[i], " ")
		}
		out()
	}
}
