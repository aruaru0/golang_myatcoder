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
	N, M := getI(), getI()

	node := make([][]int, N)
	for i := 0; i < M; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	W := getI()
	S := getStrings(N)

	// 状態数は N * W
	outDeg := make([]int, N*W)
	revGraph := make([][]int, N*W)

	// グラフの構築
	for u := 0; u < N; u++ {
		for j := 0; j < W; j++ {
			if S[u][j] == 'x' {
				continue
			}
			nj := (j + 1) % W

			// 遷移先の候補 v (自分自身、または隣接都市) を確認
			// 同じ都市に留まる場合
			if S[u][nj] == 'o' {
				outDeg[u*W+j]++
				revGraph[u*W+nj] = append(revGraph[u*W+nj], u*W+j)
			}
			// 隣接都市に移動する場合
			for _, v := range node[u] {
				if S[v][nj] == 'o' {
					outDeg[u*W+j]++
					revGraph[v*W+nj] = append(revGraph[v*W+nj], u*W+j)
				}
			}
		}
	}

	// 出次数が0 (行き止まり) の状態をキューに追加
	queue := make([]int, 0)
	for u := 0; u < N; u++ {
		for j := 0; j < W; j++ {
			if S[u][j] == 'o' && outDeg[u*W+j] == 0 {
				queue = append(queue, u*W+j)
			}
		}
	}

	// 後退解析: 行き止まりに繋がるパスを削っていく
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, prev := range revGraph[curr] {
			outDeg[prev]--
			if outDeg[prev] == 0 {
				queue = append(queue, prev)
			}
		}
	}

	// 曜日0の有効な状態の中で、削られずに残っているものがあるか判定
	ans := false
	for u := 0; u < N; u++ {
		if S[u][0] == 'o' && outDeg[u*W+0] > 0 {
			ans = true
			break
		}
	}

	if ans {
		out("Yes")
	} else {
		out("No")
	}
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
