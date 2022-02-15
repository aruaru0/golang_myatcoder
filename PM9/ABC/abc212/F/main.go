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

type pair struct{ x, y int }

func lowerBoundPair(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x || (a[i].x == x.x && a[i].y >= x.y)
	})
	if idx < len(a) && a[idx].x >= x.x { // 以上を満たせば返す
		return idx
	}
	// 以上でない場合はエラー
	return -1
}
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, Q := getI(), getI(), getI()
	A := make([]int, M)
	B := make([]int, M)
	S := make([]int, M)
	T := make([]int, M)
	bus := make([][]pair, N+1)
	// bus[i] : i番目の都市を出発するバスの出発時間とインデクスのペア
	for i := 0; i < M; i++ {
		A[i], B[i], S[i], T[i] = getI(), getI(), getI(), getI()
		bus[A[i]] = append(bus[A[i]], pair{S[i], i})
	}

	// バスの時間順にソートする
	for _, p := range bus {
		sort.Slice(p, func(i, j int) bool {
			return p[i].x < p[j].x || (p[i].x == p[j].x && p[i].y < p[j].y)
		})
	}

	// ダブリング
	table := make([][]int, 20)
	for i := range table {
		table[i] = make([]int, M)
	}
	// 初期値を設定
	for i := 0; i < M; i++ {
		// B[i]を出発するT[i]以降のバスを調べる
		itr := lowerBoundPair(bus[B[i]], pair{T[i], -1})
		if itr == -1 { // 無ければ都市にとどまる
			table[0][i] = i
		} else { // あれば、移動する
			table[0][i] = bus[B[i]][itr].y
		}
	}

	// ダブリング処理
	for i := 1; i < 20; i++ {
		for j := 0; j < M; j++ {
			table[i][j] = table[i-1][table[i-1][j]]
		}
	}

	// Queryに解答
	// 結構条件分岐が面倒（整理する必要あり）
	for l := 0; l < Q; l++ {
		x, y, z := getI(), getI(), getI()
		// 時刻x以降のy出発のバスを探す
		itr := lowerBoundPair(bus[y], pair{x, -1})
		if itr == -1 { // ない場合
			out(y)
			continue
		}
		// 現在の場所をnowに入れる
		now := bus[y][itr].y
		if z <= S[now] {
			out(y)
			continue
		}
		if z <= T[now] {
			out(A[now], B[now])
			continue
		}
		// テーブルを探索
		for i := 19; i >= 0; i-- {
			next := table[i][now]
			if T[next] < z {
				now = next
			}
		}
		itr = lowerBoundPair(bus[B[now]], pair{T[now], -1})
		if itr == -1 {
			out(B[now])
			continue
		}
		next := bus[B[now]][itr].y
		if z <= S[next] {
			out(B[now])
			continue
		}
		out(A[next], B[next])
	}
}
