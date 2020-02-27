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

// Node :
type Node struct {
	to []int
}

var route []int

func dfs(from, to, prev int) bool {
	if to == from {
		route = append(route, to)
		return true
	}
	ret := false
	for _, v := range node[from].to {
		if v == prev {
			continue
		}
		ret = dfs(v, to, from)
		if ret == true {
			break
		}
	}
	if ret == true {
		route = append(route, from)
	}
	return ret
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

type pair struct {
	f, t int
}

var node []Node

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	node = make([]Node, N)
	p := make([]pair, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		node[f].to = append(node[f].to, t)
		node[t].to = append(node[t].to, f)
		p[i] = pair{f, t}
	}

	out(node)
	f := dfs(0, N-1, -1)
	out(route, f)
	var fe, su int
	if f {
		k := len(route) / 2
		fe = route[k]
		su = route[k-1]
	}

	// 頂点１と頂点Ｎの間の経路を切って、UnionFind木で、
	// それぞれの数を数える作戦
	out("f", fe, "su", su)
	u := newUnionFind(N)
	for i := 0; i < N-1; i++ {
		if (p[i].f == fe && p[i].t == su) ||
			(p[i].f == su && p[i].t == fe) {
			continue
		}
		//		out("connect", p[i].f, p[i].t)
		u.unite(p[i].f, p[i].t)
	}
	fcnt := u.size(0)
	scnt := u.size(N - 1)
	if len(route)%2 == 1 {
		fcnt++
	}
	out(fcnt, scnt)

	if fcnt > scnt {
		fmt.Println("Fennec")
	} else {
		fmt.Println("Snuke")
	}

}
