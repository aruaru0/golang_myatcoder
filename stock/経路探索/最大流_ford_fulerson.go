package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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

const inf = math.MaxInt64 >> 10

func fordFulkerson(s, t, N int) int {
	flow := 0
	for {
		used = make([]bool, N)
		f := dfs(s, t, inf)
		if f == 0 {
			break
		}
		flow += f
		out(f)
	}
	return flow
}

func main() {
	sc.Split(bufio.ScanWords)

	/*
		node, edg
		from, to, cost
		:
		:

		4 5
		0 1 2
		0 2 1
		1 2 1
		1 3 1
		2 3 2
	*/
	N, M := getInt(), getInt()
	initG(N)
	for i := 0; i < M; i++ {
		f, t, c := getInt(), getInt(), getInt()
		addEdge(f, t, c)
	}

	for i, v := range G {
		out(i, v)
	}

	res := fordFulkerson(0, N-1, N)

	out("--------")
	for i, v := range G {
		out(i, v)
	}

	out(res)
}
