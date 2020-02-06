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

// UnionFind高速
type UnionFind struct {
	par []int
	r   []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.par = make([]int, N)
	u.r = make([]int, N)
	for i := 0; i < N; i++ {
		u.par[i] = i
		u.r[i] = 1
	}
	return u
}

func (p *UnionFind) root(x int) int {
	if p.par[x] == x {
		return x
	}
	return p.root(p.par[x])
}

func (p *UnionFind) unite(x, y int) {
	x = p.root(x)
	y = p.root(y)
	if p.r[x] < p.r[y] {
		p.par[x] = y
	} else {
		p.par[y] = x
		if p.r[x] == p.r[y] {
			p.r[x]++
		}
	}
}

func (p *UnionFind) same(x, y int) bool {
	return p.root(x) == p.root(y)
}

type node struct {
	to []int
}

func route(s int, n []node, pass []bool, cnt, N int, ans *int) {
	pass[s] = true
	cnt++
	out("now", s, "pass", pass, cnt)
	if cnt == N {
		*ans++
		out("find....")
	}
	for _, v := range n[s].to {
		//		out(v)
		apass := make([]bool, N)
		copy(apass, pass)
		if pass[v] == true {
			continue
		}
		route(v, n, apass, cnt, N, ans)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	n := make([]node, N)
	for i := 0; i < M; i++ {
		from, to := getInt()-1, getInt()-1
		out(from, to)
		n[from].to = append(n[from].to, to)
		n[to].to = append(n[to].to, from)
	}

	pass := make([]bool, N)
	var ans int
	route(0, n, pass, 0, N, &ans)

	fmt.Println(ans)
}
