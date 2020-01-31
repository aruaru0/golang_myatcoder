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
	sc.Split(bufio.ScanWords)
	N, K, L := getInt(), getInt(), getInt()

	u0 := newUnionFind(N)
	u1 := newUnionFind(N)
	for i := 0; i < K; i++ {
		x, y := getInt()-1, getInt()-1
		u0.unite(x, y)
	}

	for i := 0; i < L; i++ {
		x, y := getInt()-1, getInt()-1
		u1.unite(x, y)
	}

	type key struct {
		x, y int
	}
	m := make(map[key]int, N*2)

	for i := 0; i < N; i++ {
		k := key{u0.root(i), u1.root(i)}
		m[k]++
	}

	wr := bufio.NewWriter(os.Stdout)

	for i := 0; i < N; i++ {
		k := key{u0.root(i), u1.root(i)}
		if i != N-1 {
			fmt.Fprintf(wr, "%d ", m[k])
		} else {
			fmt.Fprintf(wr, "%d\n", m[k])
		}
	}
	wr.Flush()
}
