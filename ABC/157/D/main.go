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

// UnionFind（１）
type UnionFind struct {
	par []int
}

/* コンストラクタ */
func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

/* xの所属するグループを返す */
func (u UnionFind) root(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

/* xの所属するグループ と yの所属するグループ を合体する */
func (u UnionFind) unite(x, y int) {
	xr := u.root(x)
	yr := u.root(y)
	if xr == yr {
		return
	}
	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}

/* xとyが同じグループに所属するかどうかを返す */
func (u UnionFind) same(x, y int) bool {
	if u.root(x) == u.root(y) {
		return true
	}
	return false
}

/* xの所属するグループの木の大きさを返す */
func (u UnionFind) size(x int) int {
	return -u.par[u.root(x)]
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M, K := getInt(), getInt(), getInt()
	fr := make([]int, N+1)
	bl := make([]int, N+1)

	u := newUnionFind(N + 1)
	for i := 0; i < M; i++ {
		a, b := getInt(), getInt()
		u.unite(a, b)
		fr[a]++
		fr[b]++
	}

	for i := 0; i < K; i++ {
		c, d := getInt(), getInt()
		if u.same(c, d) {
			bl[c]++
			bl[d]++
		}
	}

	for i := 1; i <= N; i++ {
		max := u.size(i)
		new := max - fr[i] - bl[i] - 1
		fmt.Print(new, " ")
	}
	out()

}
