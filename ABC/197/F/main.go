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

type edge struct {
	to int
	c  byte
}

func pos(x, y, n int) int {
	return y*n + x
}

func xy(pos, n int) (int, int) {
	return pos % n, pos / n
}

const inf = int(1e15)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	node := make([][]edge, N)
	for i := 0; i < M; i++ {
		f, t, c := getI()-1, getI()-1, getS()[0]
		node[f] = append(node[f], edge{t, c})
		node[t] = append(node[t], edge{f, c})
	}

	// (i,j)ペアのノードを作成し、文字が同じ辺で接続できるならつなげる
	node2 := make([][]int, N*N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for _, e := range node[i] {
				for _, v := range node[j] {
					if e.c == v.c {
						node2[pos(i, j, N)] = append(node2[pos(i, j, N)], pos(e.to, v.to, N))
					}
				}
			}
		}
	}
	q := []int{pos(0, N-1, N)}
	dist := make([]int, N*N)
	for i := 0; i < N*N; i++ {
		dist[i] = inf
	}
	dist[pos(0, N-1, N)] = 0
	ans := inf
	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		x, y := xy(c, N)
		if x == y { // 場所が同じになったら、回文
			chmin(&ans, dist[c]*2)
		}
		for _, e := range node2[c] {
			xx, yy := xy(e, N)
			if x == yy && xx == y { // 次の場所が同じなら回文
				chmin(&ans, dist[c]*2+1)
			}
			if dist[e] != inf {
				continue
			}
			dist[e] = dist[c] + 1
			q = append(q, e)
		}
	}

	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
