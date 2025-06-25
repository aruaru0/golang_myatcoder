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

// Nodeは文字列のバージョンを表す構造体
type Node struct {
	par int
	add string
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N := getI()
	Q := getI()

	// nds: 文字列の各バージョンを格納する
	nds := make([]Node, 0, Q+1)
	// ルートノード(ID:0)を追加
	nds = append(nds, Node{-1, ""})

	// srv: サーバーが参照するノードID
	srv := 0
	// pcs: 各PCが参照するノードID
	pcs := make([]int, N+1)

	// クエリ処理
	for i := 0; i < Q; i++ {
		typ := getI()
		switch typ {
		case 1:
			// PC p <- Server
			p := getI()
			pcs[p] = srv
		case 2:
			// PC p <- PC p + s
			p := getI()
			s := getS()
			parent := pcs[p]
			// 新しいノードを作成
			nn := Node{parent, s}
			nds = append(nds, nn)
			// PC pのポインタを更新
			pcs[p] = len(nds) - 1
		case 3:
			// Server <- PC p
			p := getI()
			srv = pcs[p]
		}
	}

	// 最終的なサーバーの文字列を復元
	res := make([]string, 0)
	cur := srv
	for cur != -1 {
		node := nds[cur]
		if node.add != "" {
			res = append(res, node.add)
		}
		cur = node.par
	}

	// 逆順に集めた文字列の断片を結合
	ans := []byte{}
	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, res[i]...)
	}

	out(string(ans))
}
