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

var N, K int
var node [][]int
var dist []int

func bsf(n int) {
	q := make([]int, 0)
	q = append(q, n)
	dist[n] = 0
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		for _, e := range node[v] {
			if dist[e] != inf {
				continue
			}
			if dist[e] > dist[v]+1 {
				dist[e] = dist[v] + 1
				q = append(q, e)
			}
		}
	}
}

const inf = int(1e9)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, K = getInt(), getInt()
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		a, b := getInt()-1, getInt()-1
		node[a] = append(node[a], b)
		node[b] = append(node[b], a)
	}
	dist = make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	bsf(0)
	// out(dist)
	sort.Ints(dist)
	ans := 0
	for i := 0; i < N; i++ {
		if dist[i] == inf {
			break
		}
		if dist[i] != 0 {
			ans++
		}
		K--
		if K == 0 {
			break
		}
	}
	if K != 0 {
		out(-1)
		return
	}
	out(ans)
}
