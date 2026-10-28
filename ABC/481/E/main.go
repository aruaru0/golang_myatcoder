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

// 解説：https://youtu.be/UWbGRhF3Ozw?t=9564 を変換
// V は2次元ベクトルを表す構造体
type V struct {
	X, Y int
}

func (v V) Add(o V) V   { return V{v.X + o.X, v.Y + o.Y} }
func (v V) Sub(o V) V   { return V{v.X - o.X, v.Y - o.Y} }
func (v V) Mul(s int) V { return V{v.X * s, v.Y * s} }
func (v V) Div(s int) V { return V{v.X / s, v.Y / s} }

func (v V) Dot(o V) int   { return v.X*o.X + v.Y*o.Y }
func (v V) Cross(o V) int { return v.X*o.Y - v.Y*o.X }
func (v V) Norm2() int    { return v.X*v.X + v.Y*v.Y }
func (v V) Norm() float64 { return math.Sqrt(float64(v.Norm2())) }

func (v V) Normalize() V {
	n := int(v.Norm())
	if n == 0 {
		return V{0, 0}
	}
	return v.Div(n)
}

func (v V) Rotate90() V { return V{v.Y, -v.X} }

// Ort は象限（偏角ソート用）を判定
func (v V) Ort() int {
	if v.Y > 0 {
		if v.X > 0 {
			return 1
		}
		return 2
	}
	if v.X > 0 {
		return 4
	}
	return 3
}

// Less は偏角ソート用の比較関数
func (v V) Less(o V) bool {
	vo, oo := v.Ort(), o.Ort()
	if vo != oo {
		return vo < oo
	}
	return v.Cross(o) > 0
}

// Line は2点を通る直線・線分を表す構造体
type Line struct {
	S, T V
}

func (l Line) Dir() V       { return l.T.Sub(l.S) }
func (l Line) Normalize() V { return l.Dir().Normalize() }
func (l Line) Norm() int    { return int(l.Dir().Norm()) }

/*
CCW の戻り値:
+1: 反時計回り (s-t に対し s-p が反時計回り)
-1: 時計回り (s-t に対し s-p が時計回り)
+2: p-s-t (同一直線上で s の反対側)
-2: s-t-p (同一直線上で t の先)

	0: s-p-t (線分 s-t 上)
*/
func (l Line) CCW(p V) int {
	dir := l.Dir()
	ps := p.Sub(l.S)
	cr := dir.Cross(ps)
	if cr > 0 {
		return 1
	}
	if cr < 0 {
		return -1
	}
	if dir.Dot(ps) < 0 {
		return 2
	}
	if dir.Dot(l.T.Sub(p)) < 0 {
		return -2
	}
	return 0
}

func (l Line) Touch(o Line) bool {
	a := l.CCW(o.S) * l.CCW(o.T)
	b := o.CCW(l.S) * o.CCW(l.T)
	return a == 0 || b == 0 || (a == -1 && b == -1)
}

func solve() {
	n := getI()
	ps := make([]V, n)
	for i := 0; i < n; i++ {
		ps[i] = V{X: getI(), Y: getI()}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Less(ps[j])
	})

	for i := 0; i < n; i++ {
		ps = append(ps, ps[i])
	}

	ans := 0
	l := 0
	for r := 0; r < n*2; r++ {
		for r-l >= n || (Line{S: V{0, 0}, T: ps[r]}.CCW(ps[l]) > 0) {
			l++
		}
		if cur := r - l + 1; cur > ans {
			ans = cur
		}
	}

	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		solve()
	}
}
