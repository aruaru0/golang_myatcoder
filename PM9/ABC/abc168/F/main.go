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

type pair struct{ a, b int }

var dx, dy = []int{0, 1, 0, -1}, []int{1, 0, -1, 0}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	a, b, c := make([]int, n), make([]int, n), make([]int, n)
	d, e, f := make([]int, m), make([]int, m), make([]int, m)
	x := map[int]int{}
	y := map[int]int{}
	// 座標データの読み込み
	for i := 0; i < n; i++ {
		a[i], b[i], c[i] = getI(), getI(), getI()
		x[a[i]], x[b[i]], y[c[i]] = 0, 0, 0
	}
	for i := 0; i < m; i++ {
		d[i], e[i], f[i] = getI(), getI(), getI()
		x[d[i]], y[e[i]], y[f[i]] = 0, 0, 0
	}
	x[0], y[0] = 0, 0

	// 存在する座標を列挙
	xs, ys := []int{}, []int{}
	for i := range x {
		xs = append(xs, i)
	}
	for i := range y {
		ys = append(ys, i)
	}

	// ソート
	sort.Ints(xs)
	sort.Ints(ys)

	// 各座標にインデックスを入れる
	for i := 0; i < len(xs); i++ {
		x[xs[i]] = i
	}
	for i := 0; i < len(ys); i++ {
		y[ys[i]] = i
	}

	// 倍の幅の配列を用意する（マスと線の部分があるので２倍になることに注意）
	h, w := len(xs)*2, len(ys)*2
	g := make([][]bool, h)
	done := make([][]bool, h)
	for i := 0; i < h; i++ {
		g[i] = make([]bool, w)
		done[i] = make([]bool, w)
	}

	// 壁を設置
	for i := 0; i < n; i++ {
		aa, bb, cc := x[a[i]]<<1, x[b[i]]<<1, y[c[i]]<<1
		for j := aa; j <= bb; j++ {
			g[j][cc] = true
		}
	}

	for i := 0; i < m; i++ {
		dd, ee, ff := x[d[i]]<<1, y[e[i]]<<1, y[f[i]]<<1
		for j := ee; j <= ff; j++ {
			g[dd][j] = true
		}
	}

	// bsf
	q := []pair{pair{x[0] << 1, y[0] << 1}}
	done[x[0]<<1][y[0]<<1] = true
	for len(q) != 0 {
		xx, yy := q[0].a, q[0].b
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := xx+dx[i], yy+dy[i]
			if nx < 0 || nx >= h || ny < 0 || ny >= w {
				continue
			}
			if g[nx][ny] || done[nx][ny] {
				continue
			}
			done[nx][ny] = true
			q = append(q, pair{nx, ny})
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !done[i][j] {
				continue
			}
			if i == 0 || i == h-1 || j == 0 || j == w-1 {
				fmt.Fprintln(wr, "INF")
				return
			}
			if i&1 == 0 || j&1 == 0 {
				continue
			}
			ans += (xs[i>>1+1] - xs[i>>1]) * (ys[j>>1+1] - ys[j>>1])
		}
	}
	fmt.Fprintln(wr, ans)
}
