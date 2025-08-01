package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	n            int
}

// newDsu :
func newDsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
	}
	return &d
}

// Merge :
func (d DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

// Same :
func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader :
func (d DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

// Size :
func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

// Groups : original implement
func (d DSU) Groups() [][]int {
	m := make(map[int][]int)
	for i := 0; i < d.n; i++ {
		x := d.Leader(i)
		if x < 0 {
			m[i] = append(m[i], i)
		} else {
			m[x] = append(m[x], i)
		}
	}
	ret := make([][]int, len(m))
	idx := 0
	for _, e := range m {
		ret[idx] = make([]int, len(e))
		copy(ret[idx], e)
		idx++
	}
	return ret
}

type pair struct {
	u, v int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	const N = int(3e+5)
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	u := make([]int, N)
	v := make([]int, N)
	p := make([][]int, N) // 現在所属している頂点のグループ
	pRev := make([]int, N)
	e := make([]map[int]bool, N) // u->vの辺を管理

	for i := 0; i < n; i++ {
		pRev[i] = i
		p[i] = []int{i}
		e[i] = make(map[int]bool)
	}

	for i := 0; i < m; i++ {
		u[i] = getI() - 1
		v[i] = getI() - 1
		e[u[i]][v[i]] = true
		e[v[i]][u[i]] = true
	}

	q := getI()

	for i := 0; i < q; i++ {
		x := getI() - 1
		vx := pRev[u[x]]
		vy := pRev[v[x]]

		if vx != vy {
			// 辺の数　＋　頂点数
			valx := len(e[vx]) + len(p[vx])
			valy := len(e[vy]) + len(p[vy])
			// 小さい方をvxにする
			if valx > valy {
				vx, vy = vy, vx
			}

			// vxの頂点をvyに移し、頂点の親をvyに設定する
			for _, node := range p[vx] {
				p[vy] = append(p[vy], node)
				pRev[node] = vy
			}
			p[vx] = nil

			// vxの辺に関する処理
			for vz := range e[vx] {
				if vz == vy { // vyへの接続は減らして、マップから削除
					m--
					delete(e[vy], vx)
				} else {
					if _, ok := e[vy][vz]; ok { // 辺があれば辺を削除
						m--
					} else { // 辺がなければ追加（移動なので辺は増えない）
						e[vy][vz] = true
						e[vz][vy] = true
					}
					delete(e[vz], vx)
				}
			}
			e[vx] = nil // vxのマップを空にする（メモリ削減）
		}
		out(m) // 現在の辺数を表示
	}
}
