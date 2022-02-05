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

	// 先頭と末尾を取り出す
	sa := make([]string, m)
	sb := make([]string, m)
	for i := 0; i < m; i++ {
		s := getS()
		sa[i] = s[:3]
		sb[i] = s[len(s)-3:]
	}

	// 先頭と末尾をmapに追加する
	mp := make(map[string]int)
	for i := 0; i < m; i++ {
		mp[sa[i]], mp[sb[i]] = 0, 0
	}
	// インデックスをつける
	n := 0
	for e := range mp {
		mp[e] = n
		n++
	}

	//　グラフを生成する（逆向きでグラフを作る）
	to := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < m; i++ {
		a, b := mp[sa[i]], mp[sb[i]]
		to[b] = append(to[b], a)
		deg[a]++
	}

	// 端点をキューに入れる
	// ansは勝ち1,負け-1,引き分け0
	q := make([]int, 0)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			ans[i] = -1
			q = append(q, i)
		}
	}

	// 後方解析？とかいうらしい
	// 端点がなくなるまで
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		// もし、vが負け状態なら
		if ans[v] == -1 {
			// ansが0の場合は勝ちにする
			// ※１つでも相手が負け状態に遷移できるなら、勝てる
			for _, u := range to[v] {
				if ans[u] != 0 {
					continue
				}
				ans[u] = 1
				q = append(q, u)
			}
		} else { //　勝ち状態なら
			for _, u := range to[v] {
				if ans[u] != 0 {
					continue
				}
				// どこに行っても勝状態なら負けにする
				deg[u]--
				if deg[u] == 0 {
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
