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

	// --- 1. 高橋君と青木君の情報を整理 ---
	// 移動ベクトル
	vx1, vy1 := gx1-sx1, gy1-sy1
	vx2, vy2 := gx2-sx2, gy2-sy2

	// 移動時間（=距離）
	d1 := math.Sqrt(vx1*vx1 + vy1*vy1)
	d2 := math.Sqrt(vx2*vx2 + vy2*vy2)

	// 速度ベクトル（正規化された方向ベクトル）
	if d1 > 0 {
		vx1 /= d1
		vy1 /= d1
	}
	if d2 > 0 {
		vx2 /= d2
		vy2 /= d2
	}

	// --- 2. 距離を計算する関数 ---
	// 時刻tにおける2人の距離を計算する
	calcDist := func(t float64) float64 {
		t1 := math.Min(t, d1) // 高橋君の移動時間
		t2 := math.Min(t, d2) // 青木君の移動時間

		px1 := sx1 + vx1*t1
		py1 := sy1 + vy1*t1
		px2 := sx2 + vx2*t2
		py2 := sy2 + vy2*t2

		return math.Sqrt(math.Pow(px1-px2, 2) + math.Pow(py1-py2, 2))
	}

	// --- 3. 最小値候補をチェック ---
	ans := calcDist(0.0) // t=0 を初期値とする

	// 最小値候補を求める関数
	// f(t) = a*t^2 + b*t + c
	// minTime <= t <= maxTime での最小値を探索し、ansを更新
	check := func(a, b, c, minTime, maxTime float64) {
		// 端点での距離
		ans = math.Min(ans, calcDist(minTime))
		ans = math.Min(ans, calcDist(maxTime))

		// 2次関数の軸
		if a > 1e-9 { // aがほぼ0の場合は直線なので、端点のみでOK
			t_axis := -b / (2 * a)
			if t_axis > minTime && t_axis < maxTime {
				ans = math.Min(ans, calcDist(t_axis))
			}
		}
	}

	// 区間1: 0 <= t <= min(d1, d2) (両方移動中)
	// 相対位置: (sx1-sx2 + t(vx1-vx2)), (sy1-sy2 + t(vy1-vy2))
	// 距離の2乗: ((sx1-sx2) + t(vx1-vx2))^2 + ((sy1-sy2) + t(vy1-vy2))^2
	//       = (vx1-vx2)^2 + (vy1-vy2)^2 * t^2
	//       + 2 * ((sx1-sx2)(vx1-vx2) + (sy1-sy2)(vy1-vy2)) * t
	//       + (sx1-sx2)^2 + (sy1-sy2)^2
	s_dx, s_dy := sx1-sx2, sy1-sy2
	v_dx, v_dy := vx1-vx2, vy1-vy2
	a1 := v_dx*v_dx + v_dy*v_dy
	b1 := 2 * (s_dx*v_dx + s_dy*v_dy)
	c1 := s_dx*s_dx + s_dy*s_dy
	check(a1, b1, c1, 0, math.Min(d1, d2))

	// 区間2: min(d1, d2) < t <= max(d1, d2) (片方移動中)
	// d1 < d2 の場合 (高橋君が先に停止)
	if d1 < d2 {
		// 高橋君はゴール地点(gx1, gy1)にいる
		// 青木君の位置は (sx2 + vx2*t, sy2 + vy2*t)
		// 相対位置: (gx1 - (sx2+vx2*t)), (gy1 - (sy2+vy2*t))
		s_dx, s_dy = gx1-sx2, gy1-sy2
		v_dx, v_dy = -vx2, -vy2
		a2 := v_dx*v_dx + v_dy*v_dy
		b2 := 2 * (s_dx*v_dx + s_dy*v_dy)
		c2 := s_dx*s_dx + s_dy*s_dy
		check(a2, b2, c2, d1, d2)
	}
	// d2 < d1 の場合 (青木君が先に停止)
	if d2 < d1 {
		s_dx, s_dy = sx1-gx2, sy1-gy2
		v_dx, v_dy = vx1, vy1
		a2 := v_dx*v_dx + v_dy*v_dy
		b2 := 2 * (s_dx*v_dx + s_dy*v_dy)
		c2 := s_dx*s_dx + s_dy*s_dy
		check(a2, b2, c2, d2, d1)
	}

	// 区間3: max(d1, d2) < t (両方停止)
	// この区間の距離は常にゴール地点間の距離
	ans = math.Min(ans, calcDist(math.Max(d1, d2)+1.0))

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
