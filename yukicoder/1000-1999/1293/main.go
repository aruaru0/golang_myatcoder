package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
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

// Dsu :
func Dsu(n int) *DSU {
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, D, W := getI(), getI(), getI()

	uf0 := Dsu(N)
	uf1 := Dsu(N)

	// 自動車道路と歩道を別々に接続確認
	for i := 0; i < D; i++ {
		a, b := getI()-1, getI()-1
		uf0.Merge(a, b)
	}
	for i := 0; i < W; i++ {
		a, b := getI()-1, getI()-1
		uf1.Merge(a, b)
	}

	// それぞれのグループを調べる
	x := uf0.Groups()
	y := uf1.Groups()
	// それぞれの都市がどの歩道とつながるかを格納
	which := make([]int, N)
	for i := 0; i < len(y); i++ {
		for _, e := range y[i] {
			which[e] = i
		}
	}
	// out(x, y)
	// out(which)

	ans := 0
	for i := 0; i < len(x); i++ {
		n := 0
		used := make([]bool, len(y))
		for _, e := range x[i] {
			// 都市の自動車道路でつながるグループとつながる
			// 歩道の数を計算（同じ歩道グループに他の都市から
			// つながっている場合は、重複計算しないように
			if used[which[e]] == false {
				n += len(y[which[e]]) - 1
				used[which[e]] = true
			} else {
				// 複数箇所からつながっている場合は、引く
				n--
			}
		}
		m := len(x[i])
		// 自動車道路でつながっている数　＋　歩道で行ける数
		ans += m*(m-1) + m*n
		// out(x[i], n, m)
	}
	out(ans)
}
