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

// [距離, インデックス] のペア
type distInfo [2]int

// 木DPをして、それぞれのノードからの最大距離と、その最大距離を持つノードのインデックスを計算する
// (同じ距離があればインデックスが大きい方)
func solveTreeDistances(N int, node [][]int) ([]int, []int) {
	// dp1: 各ノードにおける部分木内の最長距離とインデックス
	dp1 := make([]distInfo, N)
	// dp2: 親から子へ再帰時の距離（全方位木DP）
	dp2 := make([]distInfo, N)

	ansDist := make([]int, N)
	ansIdx := make([]int, N)

	// 最初のDFS: 各部分木の最長距離とインデックスを計算
	var dfs1 func(v, parent int)
	dfs1 = func(v, parent int) {
		dp1[v] = distInfo{0, v} // [distance, index] (自分自身)
		for _, to := range node[v] {
			if to == parent {
				continue
			}
			dfs1(to, v)
			dist := dp1[to][0] + 1
			idx := dp1[to][1]

			if dp1[v][0] < dist { // より長い距離が見つかった
				dp1[v][0] = dist
				dp1[v][1] = idx
			} else if dp1[v][0] == dist { // 同じ距離ならインデックスが大きい方
				dp1[v][1] = max(dp1[v][1], idx)
			}
		}
	}

	// 2回目のDFS: 親→子に情報を伝搬（全方位DP）
	var dfs2 func(v, parent int)
	dfs2 = func(v, parent int) {
		// v の子供たちの (dp1[to][0]+1, dp1[to][1]) のペアで
		// (dist, idx) が最大のペア (first) と 2番目のペア (second) を記録

		first := distInfo{-1, -1} // 距離-1で初期化 (存在しない)
		second := distInfo{-1, -1}

		for _, to := range node[v] {
			if to == parent {
				continue
			}
			info := distInfo{dp1[to][0] + 1, dp1[to][1]}

			// 距離で比較、同じならインデックスで比較
			if info[0] > first[0] || (info[0] == first[0] && info[1] > first[1]) {
				second = first
				first = info
			} else if info[0] > second[0] || (info[0] == second[0] && info[1] > second[1]) {
				second = info
			}
		}

		for _, to := range node[v] {
			if to == parent {
				continue
			}

			// --- to に伝播する情報を計算 ---

			// 1. 親 v の親方向からの情報 (dp2[v])
			// 距離は+1, インデックスはそのまま
			parentInfo := distInfo{dp2[v][0] + 1, dp2[v][1]}

			// 2. 親 v の兄弟(to以外)の部分木からの情報
			siblingInfo := first

			// もし to の情報が first (1番目) だった場合、second (2番目) を使う
			// (dp1[to] から計算した情報が first と一致するか)
			if first[0] == dp1[to][0]+1 && first[1] == dp1[to][1] {
				siblingInfo = second
			}

			// siblingInfo の距離を v 経由 (+1) にする
			// (siblingInfo は dp1[sibling][0]+1 を元にしているので、さらに+1)
			if siblingInfo[0] != -1 {
				siblingInfo[0]++
			}

			// dp2[to] をセット (parentInfo と siblingInfo のうち良い方)
			if parentInfo[0] > siblingInfo[0] {
				dp2[to] = parentInfo
			} else if parentInfo[0] < siblingInfo[0] {
				dp2[to] = siblingInfo
			} else { // 距離が同じ
				dp2[to] = distInfo{parentInfo[0], max(parentInfo[1], siblingInfo[1])}
			}

			// 子ノードへ再帰
			dfs2(to, v)
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)

	for i := 0; i < N; i++ {
		info1 := dp1[i] // 部分木内の情報
		info2 := dp2[i] // 部分木外の情報

		if info1[0] > info2[0] {
			ansDist[i] = info1[0]
			ansIdx[i] = info1[1]
		} else if info1[0] < info2[0] {
			ansDist[i] = info2[0]
			ansIdx[i] = info2[1]
		} else { // info1[0] == info2[0] (距離が同じ)
			ansDist[i] = info1[0]
			ansIdx[i] = max(info1[1], info2[1]) // インデックスが大きい方
		}
	}
	return ansDist, ansIdx
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node := make([][]int, N)
	for i := 0; i < N-1; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	_, ansIdx := solveTreeDistances(N, node)

	for _, e := range ansIdx {
		out(e + 1)
	}
}
