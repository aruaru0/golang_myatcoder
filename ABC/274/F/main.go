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

const inf = int(1e18)

type tst struct {
	f float64
	i int
	t int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	ans := 0
	N, A := getI(), getI()
	W := make([]int, N)
	X := make([]int, N)
	V := make([]int, N)
	for i := 0; i < N; i++ {
		W[i], X[i], V[i] = getI(), getI(), getI()
	}
	calc := func(i, j int) (float64, float64) {
		if V[i] == V[j] { // 速度が同じなら
			if X[i] <= X[j] && X[j] <= X[i]+A { // 範囲内なら
				return 0, float64(inf) // ずっと入っている
			} else { // 入っていない
				return -2, -1
			}
		}
		// t1, t2を計算
		t1 := -float64(X[i]-X[j]) / float64(V[i]-V[j])
		t2 := -float64(X[i]+A-X[j]) / float64(V[i]-V[j])
		// 逆なら入れ替え
		if t1 > t2 {
			t1, t2 = t2, t1
		}
		// xs[i]-xs[i]+Aに入っている時間を返す
		return t1, t2
	}

	for i := 0; i < N; i++ { // 基準をi番目にしてループ
		ts := []tst{}
		// x[i] - x[i]+Aに入っている時間を列挙
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			t1, t2 := calc(i, j)
			ts = append(ts, tst{t1, j, 0})
			ts = append(ts, tst{t2, j, 1})
		}
		sort.Slice(ts, func(i, j int) bool {
			return ts[i].f < ts[j].f
		})
		//		out(i, ts)
		t := -float64(inf)
		tw := 0
		rm := 0
		ans = max(ans, W[i]) // i番は必ず含まれる
		for _, v := range ts {
			if t != v.f { // 同じ時刻でなければ初期化
				tw -= rm
				rm = 0
			}
			t = v.f
			if v.t == 0 { // 追加なら重さを増やす
				tw += W[v.i]
			} else { // 削除なら重さを減らす
				rm += W[v.i]
			}
			if t >= 0 {
				//				out(t, tw)
				ans = max(ans, tw+W[i])
			}
		}
	}
	out(ans)
}
