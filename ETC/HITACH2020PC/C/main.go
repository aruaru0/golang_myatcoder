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

type node struct {
	to []int
}

var g0 []int
var g1 []int
var used []int

func dfs(cnt, i, N int, n []node) {
	used[i] = 1
	if cnt%2 == 0 {
		g0 = append(g0, i)
	} else {
		g1 = append(g1, i)
	}
	for _, v := range n[i].to {
		if used[v] == 1 {
			continue
		}
		dfs(cnt+1, v, N, n)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	n := make([]node, N)
	used = make([]int, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		n[f].to = append(n[f].to, t)
		n[t].to = append(n[t].to, f)
	}

	dfs(0, 0, N, n)

	out(g0, g1)
	p := make([]int, 0, N)
	cnt0 := 0
	for i := 1; i <= N; i++ {
		if i%3 == 1 {
			p = append(p, i)
			cnt0++
		}
	}
	start1 := cnt0
	cnt1 := 0
	for i := 1; i <= N; i++ {
		if i%3 == 2 {
			p = append(p, i)
			cnt1++
		}
	}
	start2 := start1 + cnt1
	cnt3 := 0
	for i := 1; i <= N; i++ {
		if i%3 == 0 {
			p = append(p, i)
			cnt3++
		}
	}
	out("pass", len(p))

	x := make([]int, N)
	if len(g0) <= cnt3 {
		out("pat0")
		for i := 0; i < len(g0); i++ {
			x[g0[i]] = p[len(p)-1-i]
		}
		for i := 0; i < len(g1); i++ {
			x[g1[i]] = p[i]
		}
	} else if len(g1) <= cnt3 {
		out("pat1")
		for i := 0; i < len(g1); i++ {
			x[g1[i]] = p[len(p)-1-i]
		}
		for i := 0; i < len(g0); i++ {
			x[g0[i]] = p[i]
		}
	} else {
		out("pat2")
		idx := 0
		for i := 0; i < cnt0; i++ {
			x[g0[idx]] = p[i]
			idx++
		}
		out("--pass")
		for i, j := idx, 0; i < len(g0); i++ {
			out(start2 + j)
			x[g0[i]] = p[start2+j]
			j++
		}
		idx = 0
		out("--pass")
		for i := 0; i < cnt1; i++ {
			x[g1[idx]] = p[start1+i]
			idx++
		}
		out("--pass")
		for i, j := idx, len(p)-1; i < len(g1); i++ {
			out(start2+j, start1, cnt0, cnt1, cnt3)
			x[g1[i]] = p[j]
			j--
		}
		out("--pass")

	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < N; i++ {
		fmt.Fprintf(w, "%d ", x[i])
	}
	fmt.Fprintf(w, "\n")
}
