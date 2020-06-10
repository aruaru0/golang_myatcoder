package main

import (
	"bufio"
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

type Edge struct {
	from, to, cost int
}

type Route struct {
	path []int
}

// ベルマンフォード法
// SからEで、Eが関連するネガティブループを見つける
func bellmanFord(N, P, S, E int, edges []Edge) ([]int, bool) {
	d := make([]int, N+1)
	for i := 0; i <= N; i++ {
		d[i] = math.MaxInt64
	}
	d[S] = 0
	for i := 0; i < N-1; i++ {
		for _, e := range edges {
			if d[e.from] != math.MaxInt64 &&
				d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
			}
		}
	}

	negative := make([]bool, N+1)
	for i := 0; i < N; i++ {
		for _, e := range edges {
			if d[e.from] != math.MaxInt64 &&
				d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
				negative[e.to] = true
			}
			negative[e.to] = negative[e.to] || negative[e.from]
		}
	}

	return d, negative[E]
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M, P := getInt(), getInt(), getInt()
	edges := make([]Edge, M)
	for i := 0; i < M; i++ {
		f, t, c := getInt(), getInt(), getInt()
		edges[i] = Edge{f, t, -c + P}
	}

	// out(edges)
	// ここから
	d, negativeLoop := bellmanFord(N, P, 1, edges)

	if negativeLoop {
		out(-1)
	} else {
		out(max(0, -d[N]))
	}
}
