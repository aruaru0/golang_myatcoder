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

type pair struct {
	f, s int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m, q := getI(), getI(), getI()

	// グラフの読み込み
	to := make([][]int, n)
	for i := 0; i < m; i++ {
		a, b := getI()-1, getI()-1
		to[a] = append(to[a], b)
		to[b] = append(to[b], a)
	}

	// 配列を初期化
	memo := make([]pair, n) //　時刻fに値sを書き込んだ
	last := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = pair{-1, -1}
		last[i] = -1
	}

	// 解法：平方分割
	// 初期値を書き込み
	// 連結数の数により２グループに分ける big: true/false
	D := 600 // sqrt(2M)付近の値
	big := make([]bool, n)
	for i := 0; i < n; i++ {
		big[i] = len(to[i]) >= D
	}

	// big側を処理
	// toの先もbigのものだけ抜き出す
	for i := 0; i < n; i++ {
		if big[i] {
			nto := make([]int, 0)
			for _, j := range to[i] {
				if big[j] {
					nto = append(nto, j)
				}
			}
			to[i] = nto
		}
	}

	// x[i]を初期化
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = i + 1
	}

	// get関数
	get := func(v int) {
		for _, u := range to[v] {
			if big[u] && memo[u].f > last[v] {
				last[v] = memo[u].f // 最後に書き込んだ時刻を更新
				x[v] = memo[u].s    // 書き込んだ値を更新
			}
		}
	}

	// ここからメイン
	for qi := 0; qi < q; qi++ {
		v := getI() - 1 // xiを読み出し
		if big[v] {
			// big側はメモしておく
			for _, u := range to[v] {
				x[u] = x[v]
			}
			memo[v] = pair{qi, x[v]}
		} else {
			// small側は処理する
			get(v) // 隣接を見ながら処理する
			for _, u := range to[v] {
				x[u] = x[v]
				if !big[u] { // 最後に更新した時間を記録
					last[u] = qi
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !big[i] { // bigは更新
			get(i)
		}
		fmt.Fprint(wr, x[i], " ")
	}
	out()
}
