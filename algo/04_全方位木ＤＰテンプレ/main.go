package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//-------------------------------------------------------
// 全方位DPテンプレ
// https://null-mn.hatenablog.com/entry/2020/04/14/124151
//

// F : F(a,b)
type F func(a, b int) int

// Merge : Merge(a,b)
type Merge func(a, b int) int

type reRooting struct {
	V     int
	G     [][]int
	dp    [][]int
	f, g  F
	merge Merge
	mi    int
}

func newReRooting(n int, f, g F, merge Merge, mi int) *reRooting {
	var ret reRooting
	ret.V = n
	ret.G = make([][]int, n)
	ret.dp = make([][]int, n)
	ret.f = f
	ret.g = g
	ret.merge = merge
	ret.mi = mi
	return &ret
}

func (r *reRooting) addEdge(from, to int) {
	r.G[from] = append(r.G[from], to)
	r.G[to] = append(r.G[to], from)
}

func (r *reRooting) dfs1(v, p int) int {
	ret := r.mi
	for i, e := range r.G[v] {
		if e == p {
			continue
		}
		r.dp[v][i] = r.dfs1(e, v)
		ret = r.merge(ret, r.f(r.dp[v][i], e))
	}
	return r.g(ret, v)
}

func (r *reRooting) dfs2(v, p, fromParent int) {
	for i, e := range r.G[v] {
		if e == p {
			r.dp[v][i] = fromParent
			break
		}
	}
	pR := make([]int, len(r.G[v])+1)
	pR[len(r.G[v])] = r.mi
	for i := len(r.G[v]); i > 0; i-- {
		pR[i-1] = r.merge(pR[i], r.f(r.dp[v][i-1], r.G[v][i-1]))
	}
	//out("pR", pR)
	pL := r.mi
	for i, e := range r.G[v] {
		if e != p {
			val := r.merge(pL, pR[i+1])
			r.dfs2(e, v, r.g(val, v))
		}
		pL = r.merge(pL, r.f(r.dp[v][i], e))
	}
}

func (r *reRooting) calc(root int) {
	for i, e := range r.G {
		r.dp[i] = make([]int, len(e))
	}
	r.dfs1(root, -1)
	//out(r.dp)
	r.dfs2(root, -1, r.mi)
	//out(r.dp)
}

func (r *reRooting) solve(v int) int {
	ans := r.mi
	for i, e := range r.G[v] {
		ans = r.merge(ans, r.f(r.dp[v][i], e))
	}
	return r.g(ans, v)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func out(x ...interface{}) {
	fmt.Println(x...)
}

func main() {
	// sample : EDPC-V subtree
	// 部分木で各ノードを利用する回数
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1024)
	n, m := getInt(), getInt()
	var f = func(a, b int) int {
		return (a + 1) % m
	}
	var g = func(a, b int) int {
		return a
	}
	var merge = func(a, b int) int {
		return (a * b) % m
	}
	re := newReRooting(n, f, g, merge, 1)
	for i := 0; i < n-1; i++ {
		x, y := getInt()-1, getInt()-1
		re.addEdge(x, y)
	}
	re.calc(0)
	for i := 0; i < n; i++ {
		out(re.solve(i))
	}
}
