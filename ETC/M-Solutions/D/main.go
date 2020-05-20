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

// Priority Queue
// Item :
type Item struct {
	priority, value, index int
}

// PQ :
type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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

// End Priority Queue

func dfs(n, p int, node [][]int, c []int) int {
	ret := 0
	for _, v := range node[n] {
		if v == p {
			continue
		}
		ret += dfs(v, n, node, c) + min(c[n], c[v])
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	node := make([][]int, N)
	edge := make([][2]int, N-1)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
		edge[i] = [2]int{f, t}
	}
	c := make([]int, N)
	for i := 0; i < N; i++ {
		c[i] = getInt()
	}
	sort.Ints(c)

	idx := N - 1
	n := make([]int, N)
	m := make([]int, N-1)
	m[0] = c[idx-1]
	n[edge[0][0]] = c[idx]
	idx--
	n[edge[0][1]] = c[idx]
	idx--
	// out(n, m)
	for idx >= 0 {
		for i := 0; i < N-1; i++ {
			if m[i] != 0 {
				continue
			}
			if n[edge[i][0]] != 0 {
				m[i] = c[idx]
				n[edge[i][1]] = c[idx]
				idx--
			} else if n[edge[i][1]] != 0 {
				m[i] = c[idx]
				n[edge[i][0]] = c[idx]
				idx--
			}
		}
	}
	// out(idx)
	// out(n)

	ans := dfs(0, -1, node, n)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprintln(w, ans)
	for i := 0; i < N; i++ {
		fmt.Fprint(w, n[i], " ")
	}
	fmt.Fprintln(w)
}
