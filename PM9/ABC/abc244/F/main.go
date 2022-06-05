package main

import (
	"bufio"
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

type pair struct {
	s, v int
}

const inf int = 1e18

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	n, m := getI(), getI()
	node := make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	var dist [1 << 17][17]int
	N := 1 << n
	for i := 0; i < N; i++ {
		for j := 0; j < n; j++ {
			dist[i][j] = inf
		}
	}
	q := make([]pair, 0)
	// どれか１つを選んだ状態から開始
	for i := 0; i < n; i++ {
		dist[1<<i][i] = 1
		q = append(q, pair{1 << i, i})
	}

	// BSF
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		// つながっている系についてループ
		for _, u := range node[cur.v] {
			// 次状態は、現在の状態と次のノードの到達回数（偶奇）
			// つまり、よいパスのものを作る
			ns := cur.s ^ (1 << u)
			// もし、すでに探索していればスキップ（最短ではない）
			if dist[ns][u] < inf {
				continue
			}
			// 到達できるので長さ＋１する
			dist[ns][u] = dist[cur.s][cur.v] + 1
			q = append(q, pair{ns, u})
		}
	}

	ans := 0
	// 各パターンで到達できる値を集計する
	// 単純連結なのでinfはないはず
	for i := 1; i < N; i++ {
		mn := inf
		for j := 0; j < n; j++ {
			mn = min(mn, dist[i][j])
		}
		ans += mn
	}
	out(ans)
}
