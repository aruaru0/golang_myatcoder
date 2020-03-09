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

// UnionFind高速
type UnionFind struct {
	d []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.d = make([]int, N)
	for i := 0; i < N; i++ {
		u.d[i] = -1
	}
	return u
}

func (p *UnionFind) root(x int) int {
	if p.d[x] < 0 {
		return x
	}
	p.d[x] = p.root(p.d[x])
	return p.d[x]
}

func (p *UnionFind) unite(x, y int) bool {
	x = p.root(x)
	y = p.root(y)
	if x == y {
		return false
	}
	if p.d[x] > p.d[y] {
		x, y = y, x
	}
	p.d[x] += p.d[y]
	p.d[y] = x
	return true
}

func (p *UnionFind) same(x, y int) bool {
	return p.root(x) == p.root(y)
}

func (p *UnionFind) size(x int) int {
	return -p.d[p.root(x)]
}

type edge struct {
	a, b int
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	e := make([]edge, M)
	for i := 0; i < M; i++ {
		e[i].a, e[i].b = getInt()-1, getInt()-1
	}

	cnt := 0
	for i := 0; i < M; i++ {
		u := newUnionFind(N)
		for j := 0; j < M; j++ {
			if j == i {
				continue
			}
			u.unite(e[j].a, e[j].b)
		}
		if u.same(e[i].a, e[i].b) {
			cnt++
		}
	}
	out(M - cnt)
}
