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

type edge struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	e := make([]edge, M)
	for i := 0; i < M; i++ {
		f, t := getInt()-1, getInt()-1
		e[i] = edge{f, t}
	}
	u := newUnionFind(N)
	size := N * (N - 1) / 2
	// out(0)
	cnt := 0
	ans := make([]int, 0)
	ans = append(ans, size)
	for i := M - 1; i > 0; i-- {
		if !u.same(e[i].x, e[i].y) {
			x := u.size(e[i].x)
			y := u.size(e[i].y)
			if x > 1 {
				cnt -= x * (x - 1) / 2
			}
			if y > 1 {
				cnt -= y * (y - 1) / 2
			}
			u.unite(e[i].x, e[i].y)
			z := u.size(e[i].x)
			cnt += z * (z - 1) / 2
			// out(x, y, z)
		}
		// out(size - cnt)
		ans = append(ans, size-cnt)
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprintln(w, ans[i])
	}
}
