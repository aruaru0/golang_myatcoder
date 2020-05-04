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

type edge struct {
	to, num int
}

type node struct {
	e []edge
}

var edgeCol []int

func dfs(pos, prev, col int, n []node) {
	c := 0
	for _, v := range n[pos].e {
		if v.to == prev {
			continue
		}
		c++
		if c == col {
			c++
		}
		edgeCol[v.num] = c
		// out(v.to, pos, c)
		dfs(v.to, pos, c, n)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	n := make([]node, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		n[f].e = append(n[f].e, edge{t, i})
		n[t].e = append(n[t].e, edge{f, i})
	}
	edgeCol = make([]int, N-1)

	dfs(0, -1, -1, n)

	//	out(n)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	ma := 0
	for _, v := range edgeCol {
		ma = max(ma, v)
	}
	fmt.Fprintln(w, ma)
	for _, v := range edgeCol {
		fmt.Fprintln(w, v)
	}
}
