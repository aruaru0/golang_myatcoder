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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, M := getInt(), getInt()
	u := newUnionFind(N)
	for i := 0; i < M; i++ {
		a, b := getInt()-1, getInt()-1
		u.unite(a, b)
	}
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		r := u.root(i)
		m[r]++
	}
	out(len(m) - 1)
}
