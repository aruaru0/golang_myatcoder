package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type pair struct {
	f, t int
}

func dfs(p, x int, m map[int]pair, ans []int) {
	ans[p] = x + 1
	e := m[p]
	if e.f != -1 && ans[e.f] == 0 {
		dfs(e.f, x, m, ans)
	}
	if e.t != -1 && ans[e.t] == 0 {
		dfs(e.t, x, m, ans)
	}
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	N, Q := getInt(), getInt()

	n := make([]int, N)
	m := make(map[int]pair)
	for i := 0; i < N; i++ {
		m[i] = pair{-1, -1}
		n[i] = i
	}

	// out(m)
	// out(n, Q)
	// out("====")

	for i := 0; i < Q; i++ {
		f, t, x := getInt()-1, getInt()-1, getInt()-1
		p := m[x]
		q := m[n[t]]
		tmp := n[f]
		n[f] = m[x].f
		if n[f] != -1 {
			m[n[f]] = pair{m[n[f]].f, -1}
		}
		if n[t] != -1 {
			m[n[t]] = pair{q.f, x}
		}
		m[x] = pair{n[t], p.t}
		n[t] = tmp
		// out(f, t, x)
		// out(n, m)
		// out("-----")
	}

	ans := make([]int, N)
	for i := 0; i < N; i++ {
		st := n[i]
		if st != -1 {
			dfs(st, i, m, ans)
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < N; i++ {
		fmt.Fprintln(w, ans[i])
	}
}
