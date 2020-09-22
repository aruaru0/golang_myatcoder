package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

func small(N, X, M int) {
	a := X
	ans := X
	for i := 2; i <= N; i++ {
		a *= a
		a %= M
		ans += a
	}
	out(ans)
}

// 入力パターンの巡回部分を見つけ、巡回が見つかったら
// Ｎまでの合計値を計算する
//  ABC179 E など
type countMod struct {
	mod int
	m   []int
	pat []int
	sum int
	f   func(a, b int) int
}

// Fは累計処理関数　a+bのみ確認
func newCounter(n, mod int, f func(a, b int) int) *countMod {
	var ret countMod
	ret.mod = mod
	ret.m = make([]int, mod)
	ret.pat = make([]int, 0)
	ret.sum = 0
	ret.f = f
	return &ret
}

// vは次の値、Nは繰り返し回（トータル）
// 新しい値を登録する。巡回パターンが見つかると
// 初期部＋繰り返し部＋残りを加算した総和を返す
func (c *countMod) append(v, N int) (int, bool) {
	a := v
	if c.m[a] == 0 {
		c.m[a]++
		c.pat = append(c.pat, a)
		c.sum = c.f(c.sum, a)
		if len(c.pat) == N {
			return c.sum, true
		}
		return c.sum, false
	}
	start := 0
	ret := 0
	tot := c.sum
	for i := 0; i < len(c.pat); i++ {
		if c.pat[i] == a {
			start = i
			break
		}
		ret = c.f(ret, c.pat[i])
		tot -= c.pat[i]
	}
	N -= start
	loop := len(c.pat) - start
	n := N / loop
	ret += n * tot
	rest := N % loop
	for i := 0; i < rest; i++ {
		ret = c.f(ret, c.pat[start+i])
	}
	return ret, true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, X, M := getInt(), getInt(), getInt()

	c := newCounter(N, M, func(a, b int) int {
		return a + b
	})
	for {
		sum, flg := c.append(X, N)
		// out(sum, flg, X, N)
		if flg {
			out(sum)
			return
		}
		X *= X
		X %= M
	}
}
