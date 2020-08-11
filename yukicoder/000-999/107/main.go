package main

import (
	"bufio"
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

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	D := getInts(N)

	var dp [100000]int
	// for i := 0; i < N; i++ {
	// 	dp[i] = -1e18
	// }

	dp[0] = 100
	for i := 0; i < 1<<N; i++ {
		if dp[i] <= 0 {
			continue
		}
		cnt := 1
		for j := 0; j < N; j++ {
			if i&(1<<j) != 0 && D[j] < 0 {
				cnt++
			}
		}
		for j := 0; j < N; j++ {
			if i&(1<<j) == 0 {
				if D[j] > 0 {
					dp[i+(1<<j)] = min(cnt*100, max(dp[i+(1<<j)], dp[i]+D[j]))
				} else if dp[i]+D[j] > 0 {
					dp[i+(1<<j)] = max(dp[i+(1<<j)], dp[i]+D[j])
				}
			}
		}
	}
	out(dp[(1<<N)-1])
}
