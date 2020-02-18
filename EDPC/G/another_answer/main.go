package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
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

// Node :
type Node struct {
	from []int
	to   []int
	n    int
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()

	node := make([]Node, N)
	for i := 0; i < M; i++ {
		from, to := getInt()-1, getInt()-1
		node[from].to = append(node[from].to, to)
		node[to].from = append(node[to].from, from)
	}

	var s []int

	for i := 0; i < N; i++ {
		l := len(node[i].to)
		node[i].n = l
		if l == 0 {
			s = append(s, i)
		}
	}

	out(N, M, node)

	out(s)

	pos := 0
	for i := 0; i < N; i++ {
		out(pos)
		idx := s[pos]
		for j := 0; j < len(node[idx].from); j++ {
			f := node[idx].from[j]
			node[f].n--
			out(i, j, "f", f)
			if node[f].n == 0 {
				s = append(s, f)
			}
		}
		pos++
	}
	out(s)

	m := 0
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		idx := s[i]
		out("---", i)
		if len(node[idx].to) != 0 {
			for _, v := range node[idx].to {
				out(v)
				dp[idx] = max(dp[idx], dp[v]+1)
			}
		}
		if dp[idx] > m {
			m = dp[idx]
		}
	}

	out(dp)
	fmt.Println(m)
}
