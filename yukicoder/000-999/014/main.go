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

// 約数を列挙
func divisor(n int) []int {
	ret := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if i*i != n {
				ret = append(ret, n/i)
			}
		}
	}
	sort.Ints(ret)
	return ret
}

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM : find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a / GCD(a, b) * b

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	n := make([]int, 10010)
	pq := make([]priorityQueue, 10010)
	b := make([][]int, 10010)

	for i := 0; i < N; i++ {
		a[i] = getInt()
		n[a[i]]++
		b[a[i]] = divisor(a[i])
		for _, v := range b[a[i]] {
			heap.Push(&pq[v], pqi{a[i]})
		}
		// out(a[i], b[a[i]])
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	x := a[0]
	n[x]--
	fmt.Fprint(w, x, " ")
	for i := 1; i < N; i++ {
		// out(b[x], n[:10])
		minLCM := math.MaxInt64
		val := math.MaxInt64
		for _, v := range b[x] {
			k := pqi{-1}
			for len(pq[v]) > 0 {
				k = pq[v][0]
				if n[k.a] > 0 {
					break
				}
				heap.Pop(&pq[v])
			}
			if k.a != -1 && n[k.a] != 0 {
				lcm := LCM(x, k.a)
				// out(x, k.a, lcm)
				if minLCM > lcm {
					minLCM = lcm
					val = k.a
				} else if minLCM == lcm && val > k.a {
					minLCM = lcm
					val = k.a
				}
			}
		}
		n[val]--
		// out(val)
		fmt.Fprint(w, val, " ")
		x = val
		a[i] = val
	}
	fmt.Fprintln(w)

}
