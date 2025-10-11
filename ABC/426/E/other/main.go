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

func solve() {
	sx1, sy1, gx1, gy1 := getF(), getF(), getF(), getF()
	sx2, sy2, gx2, gy2 := getF(), getF(), getF(), getF()

	// 移動ベクトル
	vx1, vy1 := gx1-sx1, gy1-sy1
	vx2, vy2 := gx2-sx2, gy2-sy2

	// 移動時間（=距離）
	d1 := math.Sqrt(vx1*vx1 + vy1*vy1)
	d2 := math.Sqrt(vx2*vx2 + vy2*vy2)

	// 速度ベクトル（正規化）
	if d1 > 1e-9 {
		vx1 /= d1
		vy1 /= d1
	}
	if d2 > 1e-9 {
		vx2 /= d2
		vy2 /= d2
	}

	// 時刻tにおける2人の距離を計算する関数
	calcDist := func(t float64) float64 {
		t1 := math.Min(t, d1)
		t2 := math.Min(t, d2)

		px1 := sx1 + vx1*t1
		py1 := sy1 + vy1*t1
		px2 := sx2 + vx2*t2
		py2 := sy2 + vy2*t2

		return math.Sqrt(math.Pow(px1-px2, 2) + math.Pow(py1-py2, 2))
	}

	// 指定された範囲 [l, r] 内で最小値を探す三分探索関数
	ternarySearch := func(l, r float64) float64 {
		// 探索範囲が非常に小さい、またはループ回数が十分な場合は端点の小さい方を返す
		for i := 0; i < 50; i++ {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			if calcDist(m1) < calcDist(m2) {
				r = m2
			} else {
				l = m1
			}
		}
		return calcDist(l)
	}

	// --- メイン処理: 各フェーズで最小値を求める ---

	// 最小値の初期値として t=0 の距離を設定
	ans := calcDist(0)

	// フェーズ1の探索
	t_end1 := math.Min(d1, d2)
	if t_end1 > 1e-9 { // フェーズ1の区間がほぼ0でない場合のみ探索
		ans = math.Min(ans, ternarySearch(0, t_end1))
	}

	// フェーズ2の探索
	t_start2 := math.Min(d1, d2)
	t_end2 := math.Max(d1, d2)
	if t_end2-t_start2 > 1e-9 { // フェーズ2の区間がほぼ0でない場合
		ans = math.Min(ans, ternarySearch(t_start2, t_end2))
	}

	// フェーズの境界値での距離もチェック
	ans = math.Min(ans, calcDist(d1))
	ans = math.Min(ans, calcDist(d2))

	fmt.Fprintf(wr, "%.15f\n", ans)
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
