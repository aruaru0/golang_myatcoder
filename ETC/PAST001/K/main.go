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
	to []int
}

var begin []int
var end []int
var cnt = 0

func dfs(cur, prev int, n []node) {
	begin[cur] = cnt
	cnt++
	for _, v := range n[cur].to {
		if v == prev {
			continue
		}
		dfs(v, cur, n)
	}
	end[cur] = cnt
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	n := make([]node, N)
	begin = make([]int, N)
	end = make([]int, N)
	root := -1
	for i := 0; i < N; i++ {
		a := getInt()
		if a == -1 {
			root = i
		} else {
			a--
			n[a].to = append(n[a].to, i)
			n[i].to = append(n[i].to, a)
		}
	}

	dfs(root, -1, n)

	Q := getInt()
	for i := 0; i < Q; i++ {
		a, b := getInt()-1, getInt()-1
		if begin[b] <= begin[a] && end[a] <= end[b] {
			out("Yes")
		} else {
			out("No")
		}
	}

}
