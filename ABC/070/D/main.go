package main

import (
	"bufio"
	"fmt"
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

type edge struct {
	to, cost int
}

type node struct {
	to   []edge
	cost int
}

func dfs(from, prev, cnt int, n []node) {
	n[from].cost = cnt
	for _, v := range n[from].to {
		if v.to == prev {
			continue
		}
		dfs(v.to, from, cnt+v.cost, n)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	n := make([]node, N)
	for i := 0; i < N-1; i++ {
		f, t, c := getInt()-1, getInt()-1, getInt()
		n[f].to = append(n[f].to, edge{t, c})
		n[t].to = append(n[t].to, edge{f, c})
	}
	Q, K := getInt(), getInt()
	dfs(K-1, -1, 0, n)

	for i := 0; i < Q; i++ {
		x, y := getInt()-1, getInt()-1
		out(n[x].cost + n[y].cost)
	}
}
