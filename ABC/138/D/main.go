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

type node struct {
	to  []int
	val int
}

var used []bool

func dfs(i int, n []node) {
	used[i] = true
	for _, v := range n[i].to {
		if used[v] {
			continue
		}
		n[v].val += n[i].val
		dfs(v, n)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N, Q := getInt(), getInt()
	n := make([]node, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		n[f].to = append(n[f].to, t)
		n[t].to = append(n[t].to, f)
	}

	for i := 0; i < Q; i++ {
		p, x := getInt()-1, getInt()
		n[p].val += x
	}

	used = make([]bool, N)
	dfs(0, n)
	for _, v := range n {
		fmt.Print(v.val, " ")
	}
	out()
}
