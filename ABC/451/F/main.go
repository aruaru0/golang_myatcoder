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

type DSU struct {
	parentOrSize []int
	diffWeight   []int // 親との色の違い (0 or 1)
	cnt0         []int // 根と同じ色の頂点数
	cnt1         []int // 根と違う色の頂点数
	n            int
	isBipartite  bool // グラフ全体が二部グラフかどうか
}

func newDsu(n int) *DSU {
	d := &DSU{
		n:            n,
		parentOrSize: make([]int, n),
		diffWeight:   make([]int, n),
		cnt0:         make([]int, n),
		cnt1:         make([]int, n),
		isBipartite:  true,
	}
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
		d.diffWeight[i] = 0
		d.cnt0[i] = 1 // 最初は自分自身のみなので同色が1
		d.cnt1[i] = 0 // 異色は0
	}
	return d
}

// Leader : 経路圧縮しつつ、根との距離（色）を確定させる
func (d *DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	root := d.Leader(d.parentOrSize[a])
	// 親の親（根）までの距離を累積
	d.diffWeight[a] = (d.diffWeight[a] + d.diffWeight[d.parentOrSize[a]]) % 2
	d.parentOrSize[a] = root
	return root
}

// Weight : 根を基準とした時の色 (0 or 1) を取得
func (d *DSU) Weight(a int) int {
	d.Leader(a) // 経路圧縮を走らせて最新の距離を取得
	return d.diffWeight[a]
}

// Merge : 辺 (a, b) を追加し、矛盾があれば isBipartite を false にする
func (d *DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)

	// a と b の現在の色の関係（0:同色, 1:異色）
	wa, wb := d.Weight(a), d.Weight(b)

	if x == y {
		// 同じ色同士を繋ごうとしたら二部グラフではない
		if wa == wb {
			d.isBipartite = false
		}
		return x
	}

	// マージテク (size による最適化)
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
		wa, wb = wb, wa
	}

	// x を y の親にする。色の差は (wa + wb + 1) % 2 となる
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	d.diffWeight[y] = (wa + wb + 1) % 2

	// y を x にマージする際の頂点数の更新
	if d.diffWeight[y] == 0 {
		// y が x と同色になる場合
		d.cnt0[x] += d.cnt0[y]
		d.cnt1[x] += d.cnt1[y]
	} else {
		// y が x と異色になる場合
		d.cnt0[x] += d.cnt1[y]
		d.cnt1[x] += d.cnt0[y]
	}

	return x
}

func (d *DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// IsBipartite : 現在のグラフが二部グラフか返す
func (d *DSU) IsBipartite() bool {
	return d.isBipartite
}

// GetMinBlack : その頂点が属する連結成分において、少なく塗れる方の頂点数を返す
func (d *DSU) GetMinBlack(a int) int {
	root := d.Leader(a)
	if d.cnt0[root] < d.cnt1[root] {
		return d.cnt0[root]
	}
	return d.cnt1[root]
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N, Q := getI(), getI()
	uf := newDsu(N)

	ans := 0

	for qi := 0; qi < Q; qi++ {
		u, v := getI()-1, getI()-1

		if !uf.Same(u, v) {
			// マージする前に、別々だった頃の最小値を引く
			ans -= uf.GetMinBlack(u)
			ans -= uf.GetMinBlack(v)

			uf.Merge(u, v)

			// マージ後の新しい最小値を足す
			ans += uf.GetMinBlack(u)
		} else {
			// 既に同じ成分でも、二部グラフ判定の更新のためにMergeは呼ぶ
			uf.Merge(u, v)
		}

		if !uf.IsBipartite() {
			out(-1)
		} else {
			out(ans)
		}
	}
}
