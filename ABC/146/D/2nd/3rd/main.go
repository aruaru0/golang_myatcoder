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

var col []int
var mcol int
var used []bool

func bfs(edge, enum [][]int) {
	q := []int{0}
	for len(q) != 0 {
		v := q[0]
		used[v] = true
		q = q[1:]
		c := 1
		ng := 0
		for _, e := range enum[v] {
			if col[e] != 0 {
				ng = col[e]
				break
			}
		}
		for i, e := range edge[v] {
			if used[e] {
				continue
			}
			if c == ng {
				c++
			}
			mcol = max(mcol, c)
			col[enum[v][i]] = c
			c++
			q = append(q, e)
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	node := make([][]int, N)
	enum := make([][]int, N)
	for i := 0; i < N-1; i++ {
		a, b := getInt()-1, getInt()-1
		node[a] = append(node[a], b)
		node[b] = append(node[b], a)
		enum[a] = append(enum[a], i)
		enum[b] = append(enum[b], i)
	}
	col = make([]int, N-1)
	used = make([]bool, N)
	bfs(node, enum)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, mcol)
	for _, e := range col {
		fmt.Fprintln(w, e)
	}
}
