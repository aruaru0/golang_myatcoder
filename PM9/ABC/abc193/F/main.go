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

// 最大流を求めるプログラム　Ford-Fulkerson法
type edge struct {
	to, cap, rev int
}

type node struct {
	to []edge
}

// G :
var G []node
var used []bool

// initG : グラフの初期化
func initG(N int) {
	G = make([]node, N)
}

func addEdge(from, to, cap int) {
	G[from].to = append(G[from].to, edge{to, cap, len(G[to].to)})
	G[to].to = append(G[to].to, edge{from, 0, len(G[from].to) - 1})
}

func dfs(v, t, f int) int {
	if v == t {
		return f
	}
	used[v] = true
	for i, e := range G[v].to {
		if used[e.to] || e.cap <= 0 {
			continue
		}
		d := dfs(e.to, t, min(f, e.cap))
		if d > 0 {
			G[v].to[i].cap -= d
			G[e.to].to[e.rev].cap += d
			return d
		}
	}
	return 0
}

func fordFulkerson(s, t, N int) int {
	flow := 0
	for {
		used = make([]bool, N)
		f := dfs(s, t, inf)
		if f == 0 {
			break
		}
		flow += f
	}
	return flow
}

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	c := make([][]byte, n)
	for i := 0; i < n; i++ {
		c[i] = []byte(getS())
	}
	// 問題を簡単にするために、市松模様で白黒を反転
	// これで、同色がつながっている数を数える問題になる。
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i+j)%2 == 1 {
				switch c[i][j] {
				case 'W':
					c[i][j] = 'B'
				case 'B':
					c[i][j] = 'W'
				}
			}
		}
	}

	// 最大流（最小カット）の問題として解く
	initG(n*n + 2)
	// SとTを定義
	S, T := n*n, n*n+1
	tot := 0 // totは、総隣接数
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			idx := i*n + j // 現在の場所(i,j)
			switch c[i][j] {
			case 'W':
				// Wのときは、S側にinfでつなぐ
				addEdge(S, idx, inf)
			case 'B':
				// Bのときは、T側にinfでつなぐ
				addEdge(idx, T, inf)
			}
			if i > 0 {
				// iが0より大きければ、(i-1, j)と双方向でつなぐ
				idx2 := (i-1)*n + j
				addEdge(idx, idx2, 1)
				addEdge(idx2, idx, 1)
				tot++
			}
			if j > 0 {
				// jが0より大きければ、(i, j-1)と双方向でつなぐ
				idx2 := i*n + j - 1
				addEdge(idx, idx2, 1)
				addEdge(idx2, idx, 1)
				tot++
			}
		}
	}
	// 最小カットを求める
	cut := fordFulkerson(S, T, n*n+2)
	// 総接続数-最小カット数が、白黒の同色がつながっている数
	out(tot - cut)
}
