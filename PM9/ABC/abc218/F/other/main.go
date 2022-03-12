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

func bsf(c int, node [][]int) []int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}

	q := []int{c}
	dist[c] = 0
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, e := range node[cur] {
			if dist[e] != inf {
				continue
			}
			dist[e] = dist[cur] + 1
			q = append(q, e)
		}
	}

	return dist
}

func bsf2(c int, node [][]int, f, t int) int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}

	q := []int{c}
	dist[c] = 0
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, e := range node[cur] {
			if cur == f && e == t {
				continue
			}
			if dist[e] != inf {
				continue
			}
			dist[e] = dist[cur] + 1
			q = append(q, e)
		}
	}

	return dist[N-1]
}

const inf = int(1e10)

var N, M int

type edge struct {
	s, t int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node := make([][]int, N)
	rnode := make([][]int, N)
	e := make([]edge, M)
	for i := 0; i < M; i++ {
		s, t := getI()-1, getI()-1
		node[s] = append(node[s], t)
		rnode[t] = append(rnode[t], s)
		e[i] = edge{s, t}
	}

	// 1 --> 各点への最短距離を求める
	dist1 := bsf(0, node)
	// N --> 各点への最短距離を求める（逆方向のグラフを利用）
	distN := bsf(N-1, rnode)

	minD := dist1[N-1]
	for i := 0; i < M; i++ {
		// 辺e[i]がつなぐ辺の両端に対して1->s t->Nを求め、距離を計算
		d := dist1[e[i].s] + distN[e[i].t] + 1
		// dが最短距離の場合は、s->tを切った場合の最短距離を計算
		if d == minD {
			ret := bsf2(0, node, e[i].s, e[i].t)
			if ret == inf {
				out(-1)
			} else {
				out(ret)
			}
		} else {
			// そうでなければ、最短距離はカットされていない
			if minD == inf {
				out(-1)
			} else {
				out(minD)
			}
		}
	}

}
