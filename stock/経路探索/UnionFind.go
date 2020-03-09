package main

import "fmt"

func out(x ...interface{}) {
	fmt.Println(x...)
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

func main() {
	u := newUnionFind(6)
	u.unite(0, 1)
	u.unite(0, 2)
	u.unite(1, 3)
	u.unite(4, 5)

	out(u.root(0))
	out(u.same(0, 3))
	out(u.same(0, 4))
}
