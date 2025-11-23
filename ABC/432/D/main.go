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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
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

type Rect struct {
	lx, rx, ly, ry int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	n, x, y := getI(), getI(), getI()
	rects := make([]Rect, 0)
	rects = append(rects, Rect{0, x, 0, y})

	// 分割をシミュレーションする
	for i := 0; i < n; i++ {
		c, a, b := getS(), getI(), getI()
		tmp := make([]Rect, 0)
		for _, e := range rects {
			if c == "X" {
				if e.lx < a && a < e.rx { // aを跨いだブロック分割
					tmp = append(tmp, Rect{e.lx, a, e.ly - b, e.ry - b})
					tmp = append(tmp, Rect{a, e.rx, e.ly + b, e.ry + b})
				} else if e.rx <= a { // 小さい
					tmp = append(tmp, Rect{e.lx, e.rx, e.ly - b, e.ry - b})
				} else { // 大きい
					tmp = append(tmp, Rect{e.lx, e.rx, e.ly + b, e.ry + b})
				}
			} else {
				if e.ly < a && a < e.ry { // aを跨いだブロックは分割
					tmp = append(tmp, Rect{e.lx - b, e.rx - b, e.ly, a})
					tmp = append(tmp, Rect{e.lx + b, e.rx + b, a, e.ry})
				} else if e.ry <= a {
					tmp = append(tmp, Rect{e.lx - b, e.rx - b, e.ly, e.ry})
				} else {
					tmp = append(tmp, Rect{e.lx + b, e.rx + b, e.ly, e.ry})
				}
			}
		}
		rects = tmp
	}

	m := len(rects)
	uf := newDsu(m)
	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			a := rects[i]
			b := rects[j]
			cx := min(a.rx, b.rx) - max(a.lx, b.lx) // 重なりがあるかないか
			cy := min(a.ry, b.ry) - max(a.ly, b.ly) // 重なりがあるかないか
			if cx < 0 || cy < 0 {                   // なければ、処理しない
				continue
			}
			if cx != 0 || cy != 0 { // どちらかが0でなければ、重なっているのでマージ
				uf.Merge(i, j)
			}
		}
	}

	// エリアを計算する
	areas := make([]int, m)
	for i := 0; i < m; i++ {
		a := rects[i]
		areas[uf.Leader(i)] += (a.rx - a.lx) * (a.ry - a.ly)
	}

	sort.Slice(areas, func(i, j int) bool {
		return areas[i] > areas[j]
	})

	// ０のエリアを除去
	for i := 0; i < len(areas); i++ {
		if areas[i] == 0 {
			areas = areas[:i]
			break
		}
	}

	sort.Ints(areas)
	out(len(areas))
	outSlice(areas)
}
