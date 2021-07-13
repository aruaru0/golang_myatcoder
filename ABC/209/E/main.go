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

const n = 52 * 52 * 52

func conv(s string) int {
	tot := 0
	for i := 0; i < 3; i++ {
		var x int
		if 'a' <= s[i] && s[i] <= 'z' {
			x = int(s[i] - 'a')
		} else {
			x = int(s[i]-'A') + 26
		}
		tot = tot*52 + x
	}
	return tot
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	m := getI()

	// 先頭と末尾の３文字を取得
	sa := make([]string, m)
	sb := make([]string, m)
	for i := 0; i < m; i++ {
		s := getS()
		sa[i] = s[:3]
		sb[i] = s[len(s)-3:]
	}
	// マップに登録
	mp := make(map[string]int)
	for i := 0; i < m; i++ {
		mp[sa[i]], mp[sb[i]] = 0, 0
	}
	// ノード番号を割り当てる
	n := 0
	for e := range mp {
		mp[e] = n
		n++
	}
	// グラフを生成
	to := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < m; i++ {
		a, b := mp[sa[i]], mp[sb[i]]
		// しりとりを逆にたどるグラフを生成
		to[b] = append(to[b], a)
		deg[a]++
	}

	out(mp)
	out(to)

	q := make([]int, 0)
	ans := make([]int, n)
	// 末端の列を取得
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			ans[i] = -1 // 末端はTakahashi
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		if ans[v] == -1 { // Takahashiの勝ちの場合
			for _, u := range to[v] {
				if ans[u] != 0 {
					continue
				}
				ans[u] = 1 // 次のノードは青木の勝ち
				q = append(q, u)
			}
		} else { // そうでない場合(0または１の場合)
			for _, u := range to[v] {
				if ans[u] != 0 {
					continue
				}
				deg[u]--         // 接続数を減らす
				if deg[u] == 0 { // 接続数が０になったら
					ans[u] = -1
					q = append(q, u)
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		b := mp[sb[i]]
		switch ans[b] {
		case -1:
			out("Takahashi")
		case 0:
			out("Draw")
		case 1:
			out("Aoki")
		}
	}
}
