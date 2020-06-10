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

// Data :
type Data struct {
	x, y, id int
}

var sw bool

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	if sw {
		return p[i].y < p[j].y
	}
	return p[i].x < p[j].x
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	x := make(Datas, N)
	y := make(Datas, N)
	for i := 0; i < N; i++ {
		a, b := getInt(), getInt()
		x[i] = Data{a, b, i}
		y[i] = Data{a, b, i}
	}
	sort.Sort(x)
	sw = true
	sort.Sort(y)

	d := make(Datas, 0)
	for i := 1; i < N; i++ {
		di := x[i].x - x[i-1].x
		d = append(d, Data{di, x[i].id, x[i-1].id})
	}
	for i := 1; i < N; i++ {
		di := y[i].y - y[i-1].y
		d = append(d, Data{di, y[i].id, y[i-1].id})
	}
	sw = false
	sort.Sort(d)
	// out(d)
	u := newUnionFind(N)
	ans := 0
	for _, v := range d {
		if u.same(v.y, v.id) {
			continue
		}
		u.unite(v.y, v.id)
		ans += v.x
	}
	out(ans)
}
