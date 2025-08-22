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

// Aho-Corasick アルゴリズムの実装
// https://youtu.be/BYoRvdgI5EU?t=9633 のC++コードを参考にしています
type Aho struct {
	to   []map[rune]int // 各ノードからの遷移先を保持するマップのスライス
	cnt  []int          // 各ノードがパターン終端である回数をカウント
	mask []int          // 各ノードがどのパターンの終端であるかを示すビットマスク
	fail []int          // 失敗リンク（サフィックスリンク）
}

// NewAho は新しいAho-Corasickオートマトンを作成します
func NewAho() *Aho {
	return &Aho{
		to:   []map[rune]int{{}}, // ルートノードで初期化
		cnt:  []int{0},
		mask: []int{0},
	}
}

// Add はパターン文字列をAho-Corasickトライに挿入します。
// 'i' はパターンの識別子（マスク用）です。
func (a *Aho) Add(s string, i int) int {
	v := 0
	for _, c := range s {
		if _, ok := a.to[v][c]; !ok {
			// 新しい遷移先がない場合、ノードを新しく作成
			a.to[v][c] = len(a.to)
			a.to = append(a.to, map[rune]int{})
			a.cnt = append(a.cnt, 0)
			a.mask = append(a.mask, 0)
		}
		v = a.to[v][c]
	}
	a.cnt[v]++            // このノードで終了するパターンの数をインクリメント
	a.mask[v] |= (1 << i) // パターンの識別子でこのノードをマーク
	return v
}

// Init はBFSを使用してオートマトンの失敗リンクを構築します。
func (a *Aho) Init() {
	a.fail = make([]int, len(a.to))
	for i := range a.fail {
		a.fail[i] = -1 // 失敗リンクを-1で初期化
	}

	q := []int{0}

	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		out(v, a.to[v], a.fail[v])
		for c, u := range a.to[v] {
			// ノードuの失敗リンクを計算
			// 親vの失敗リンクから文字cで辿れる状態がuの失敗リンク
			a.fail[u] = a.nextState(a.fail[v], c)
			// 失敗リンク先のカウントとマスクを現在のノードに伝播
			a.cnt[u] += a.cnt[a.fail[u]]
			a.mask[u] |= a.mask[a.fail[u]]
			// q.PushBack(u)
			q = append(q, u)
		}
	}
}

// nextState は現在の状態 'v' と文字 'c' に基づいて次の状態を見つける、
// C++のoperator()に相当するメソッドです。
func (a *Aho) nextState(v int, c rune) int {
	for v != -1 {
		if u, ok := a.to[v][c]; ok {
			// 直接遷移できる場合
			return u
		}
		// 直接遷移できない場合、失敗リンクを辿る
		v = a.fail[v]
	}
	// 失敗リンクを辿っても一致する遷移がない場合、ルートに戻る
	return 0
}

// GetMask は状態のマスクを取得する、C++のoperator[]に相当するメソッドです。
func (a *Aho) GetMask(v int) int {
	return a.mask[v]
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, l := getI(), getI()
	S := getStrings(n)

	aho := NewAho()
	for i := 0; i < n; i++ {
		aho.Add(S[i], i)
	}

	aho.Init()
	out(aho.to)
	out(aho.cnt)
	out(aho.mask)
	out(aho.fail)

	m := len(aho.to)
	n2 := 1 << n

	const mod = 998244353
	var dp [101][1 << 8][85]int
	dp[0][0][0] = 1
	for i := 0; i < l; i++ {
		for s := 0; s < n2; s++ {
			for j := 0; j < m; j++ {
				now := dp[i][s][j]
				if now == 0 {
					continue
				}
				for c := 0; c < 26; c++ {
					ni := i + 1
					nj := aho.nextState(j, rune('a'+c))
					ns := s | aho.mask[nj]
					dp[ni][ns][nj] += now
					dp[ni][ns][nj] %= mod
				}
			}
		}
	}
	ans := 0
	for j := 0; j < m; j++ {
		ans += dp[l][n2-1][j]
		ans %= mod
	}
	out(ans)

}
