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

func dfs(n, p int, node [][]int, c []int) int {
	ret := 0
	for _, v := range node[n] {
		if v == p {
			continue
		}
		ret += dfs(v, n, node, c) + min(c[n], c[v])
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	node := make([][]int, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
	}

	c := make([]int, N)
	for i := 0; i < N; i++ {
		c[i] = getInt()
	}
	sort.Ints(c)

	n := make([]int, N)

	// BSF
	idx := N - 1
	s := make([]int, 0)
	s = append(s, 0)
	for len(s) > 0 {
		x := s[0]
		s = s[1:]
		// out(x, s, node[x])
		n[x] = c[idx]
		idx--
		for _, e := range node[x] {
			if n[e] != 0 {
				continue
			}
			s = append(s, e)
		}
	}

	ans := dfs(0, -1, node, n)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprintln(w, ans)
	for i := 0; i < N; i++ {
		fmt.Fprint(w, n[i], " ")
	}
	fmt.Fprintln(w)
}
