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

type node struct {
	to []int
}

func solve(i int, n []node) int {
	m := make(map[int]int)
	m[i] = 1
	for _, v := range n[i].to {
		m[v] = 1
	}
	k := make(map[int]int)
	for _, v := range n[i].to {
		if v == i {
			continue
		}
		for _, w := range n[v].to {
			if m[w] != 1 {
				k[w] = 1
			}
		}
	}
	return len(k)
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	n := make([]node, N)
	for i := 0; i < M; i++ {
		from, to := getInt()-1, getInt()-1
		n[from].to = append(n[from].to, to)
		n[to].to = append(n[to].to, from)
	}

	for i := 0; i < N; i++ {
		out(solve(i, n))
	}
}
