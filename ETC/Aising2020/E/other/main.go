package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ a int }

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

type dat struct {
	score, k int
}

// 解説の通り優先キューで実装
// この問題、ライブラリの充実しているＣ＋＋が優しい
// 好きでgo使っているんだけど、たまに愚痴りたい
func f() {
	N := getInt()
	ans := 0
	L := make([]dat, 0)
	R := make([]dat, 0)
	for i := 0; i < N; i++ {
		k, l, r := getInt(), getInt(), getInt()
		ans += min(l, r)
		if l > r {
			L = append(L, dat{abs(l - r), k})
		}
		if r > l && k != N {
			R = append(R, dat{abs(l - r), N - k})
		}
	}

	sort.Slice(L, func(i, j int) bool {
		return L[i].k < L[j].k
	})
	sort.Slice(R, func(i, j int) bool {
		return R[i].k < R[j].k
	})

	// out(L, R)
	// L
	pqL := priorityQueue{}
	k := 1
	for i := 0; i < len(L); i++ {
		if L[i].k != k {
			for len(pqL) > k {
				heap.Pop(&pqL)
			}
			k = L[i].k
		}
		heap.Push(&pqL, pqi{L[i].score})
	}
	for len(pqL) > k {
		heap.Pop(&pqL)
	}
	// L
	pqR := priorityQueue{}
	k = 1
	for i := 0; i < len(R); i++ {
		if R[i].k != k {
			for len(pqR) > k {
				heap.Pop(&pqR)
			}
			k = R[i].k
		}
		heap.Push(&pqR, pqi{R[i].score})
	}
	for len(pqR) > k {
		heap.Pop(&pqR)
	}
	for _, e := range pqL {
		ans += e.a
	}
	for _, e := range pqR {
		ans += e.a
	}
	// out(pqL, pqR)
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	T := getInt()
	for i := 0; i < T; i++ {
		f()
	}
}
