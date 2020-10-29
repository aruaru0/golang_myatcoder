package main

import (
	"bufio"
	"fmt"
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

const mod = int(1e9 + 7)

var cnt int

func dfs(p, n int) int {
	if len(node[p]) == 0 {
		// out(p, n, 1)
		cnt += n * 1
		cnt %= mod
		return 1
	}
	ret := 1
	for _, e := range node[p] {
		ret += dfs(e, n+1)
	}
	// out(p, n, ret)
	cnt += n * ret
	cnt %= mod
	return ret
}

var node [][]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node = make([][]int, N)
	sub := make([]int, N)
	for i := 0; i < N-1; i++ {
		from, to := getI()-1, getI()-1
		node[from] = append(node[from], to)
		sub[to]++
	}
	top := 0
	for i := 0; i < N; i++ {
		if sub[i] == 0 {
			top = i
		}
	}

	// ある辺の両端までの点の数を求め、かけ合わせた組み合わせ
	// だけ、辺が使われるので、それを累積
	dfs(top, 0)
	out(cnt)
}
