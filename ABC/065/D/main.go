package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Data :
type Data struct {
	k, v int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i].k < p[j].k
}

// Datas2 :
type Datas2 []Data2

// Data2 :
type Data2 struct {
	k, i, j, xy int
}

func (p Datas2) Len() int {
	return len(p)
}

func (p Datas2) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas2) Less(i, j int) bool {
	return p[i].k < p[j].k
}

type node struct {
	x, y int
}

// UnionFind : 高速
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

	N := getInt()
	x := make(Datas, N)
	y := make(Datas, N)
	n := make([]node, N)
	for i := 0; i < N; i++ {
		a, b := getInt(), getInt()
		x[i] = Data{a, i}
		y[i] = Data{b, i}
		n[i] = node{a, b}
	}
	// x,y でそれぞれソート
	sort.Sort(x)
	sort.Sort(y)
	out(N)
	out("x=", x)
	out("y=", y)
	out(n)

	// ｘとｙの差分列を作成
	d := make(Datas2, (N-1)*2)
	for i := 1; i < N; i++ {
		d[(i-1)*2] = Data2{x[i].k - x[i-1].k, x[i].v, x[i-1].v, 0}
		d[(i-1)*2+1] = Data2{y[i].k - y[i-1].k, y[i].v, y[i-1].v, 1}
	}
	// ソート
	sort.Sort(d)
	out(d)

	// 最小全域木を求める
	cnt := 0
	ans := 0
	u := newUnionFind(N)
	for _, v := range d {
		if u.same(v.i, v.j) == false {
			ans += v.k
			u.unite(v.i, v.j)
			cnt++
		}
		if cnt == N-1 {
			break
		}
	}
	fmt.Println(ans)
}
