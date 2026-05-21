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

// ----------------------------------------
// modint
// ----------------------------------------
type modint struct {
	mod       int
	factMemo  []int
	ifactMemo []int
}

func newModint(m int) *modint {
	var ret modint
	ret.mod = m
	ret.factMemo = []int{1, 1}
	ret.ifactMemo = []int{1, 1}
	return &ret
}

func (m *modint) add(a, b int) int {
	ret := (a + b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) sub(a, b int) int {
	ret := (a - b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) mul(a, b int) int {
	a %= m.mod
	b %= m.mod
	ret := a * b % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) div(a, b int) int {
	a %= m.mod
	ret := a * m.modinv(b)
	ret %= m.mod
	return ret
}

func (m *modint) pow(p, n int) int {
	ret := 1
	x := p % m.mod
	for n != 0 {
		if n%2 == 1 {
			ret *= x
			ret %= m.mod
		}
		n /= 2
		x = x * x % m.mod
	}
	return ret
}

// 逆元を使った割り算（MOD）
// mod. m での a の逆元 a^{-1} を計算する
// func (m *modint) modinv(a int) int {
// 	b := m.mod
// 	u := 1
// 	v := 0
// 	for b != 0 {
// 		t := a / b
// 		a -= t * b
// 		a, b = b, a
// 		u -= t * v
// 		u, v = v, u
// 	}
// 	u %= m.mod
// 	if u < 0 {
// 		u += m.mod
// 	}
// 	return u
// }

// 拡張オイラーの互除法で逆元を求める
func (mm *modint) modinv(a int) int {
	m := mm.mod
	b, u, v := m, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= m
	if u < 0 {
		u += m
	}
	return u
}

// -----------------------------------------------
// 行列累乗
// 　A[][]のp乗を求める
// -----------------------------------------------
func (m *modint) powModMatrix(A [][]int, p int) [][]int {
	N := len(A)
	ret := make([][]int, N)
	for i := 0; i < N; i++ {
		ret[i] = make([]int, N)
		ret[i][i] = 1
	}

	for p > 0 {
		if p&1 == 1 {
			ret = m.mulMod(ret, A)
		}
		A = m.mulMod(A, A)
		p >>= 1
	}

	return ret
}

func (m *modint) mulMod(A, B [][]int) [][]int {
	H := len(A)
	W := len(B[0])
	K := len(A[0])
	C := make([][]int, H)
	for i := 0; i < W; i++ {
		C[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j]
				C[i][j] %= m.mod
			}
		}
	}

	return C
}

// ---------------------------------------------------
// nCk 計算関連：　TELすることがあるかも
//
//	※pow(x, p-2)を何度も取るので
//
// 厳しそうな場合は、ここを削除して高速なのを使う
// ---------------------------------------------------
func (m *modint) mfact(n int) int {
	if len(m.factMemo) > n {
		return m.factMemo[n]
	}
	if len(m.factMemo) == 0 {
		m.factMemo = append(m.factMemo, 1)
	}
	for len(m.factMemo) <= n {
		size := len(m.factMemo)
		m.factMemo = append(m.factMemo, m.factMemo[size-1]*size%m.mod)
	}
	return m.factMemo[n]
}

func (m *modint) mifact(n int) int {
	if len(m.ifactMemo) > n {
		return m.ifactMemo[n]
	}
	if len(m.ifactMemo) == 0 {
		m.factMemo = append(m.ifactMemo, 1)
	}
	for len(m.ifactMemo) <= n {
		size := len(m.ifactMemo)
		m.ifactMemo = append(m.ifactMemo, m.ifactMemo[size-1]*m.pow(size, m.mod-2)%m.mod)
	}
	return m.ifactMemo[n]
}

func (m *modint) nCr(n, r int) int {
	if n == r {
		return 1
	}
	if n < r || r < 0 {
		return 0
	}
	ret := 1
	ret *= m.mfact(n)
	ret %= m.mod
	ret *= m.mifact(r)
	ret %= m.mod
	ret *= m.mifact(n - r)
	ret %= m.mod
	return (ret)
}

// Aho-Corasick アルゴリズムの実装
// ABC-419FのC++コードを参考にしています
type Aho struct {
	to   []map[rune]int // 各ノードからの遷移先を保持するマップのスライス
	cnt  []int          // 各ノードがパターン終端である回数をカウント
	mask []int          // 各ノードがどのパターンの終端であるかを示すビットマスク（※パターン数が64個まで）
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
	N, K := getI(), getI()
	s := getStrings(K)

	mod := newModint(998244353)

	nx := NewAho()
	for i, st := range s {
		nx.Add(st, i)
	}
	nx.Init()

	M := len(nx.to)
	A := make([][]int, M)
	for i := 0; i < M; i++ {
		A[i] = make([]int, M)
	}

	// 各状態 v から次の文字 c を追加したときの遷移を考える
	for v := 0; v < M; v++ {
		// すでに禁止文字列を含んでしまっている状態からは遷移させない
		if nx.GetMask(v) != 0 {
			continue
		}

		for c := 'a'; c <= 'z'; c++ {
			next := nx.nextState(v, c)
			// 遷移先も禁止文字列を含まない安全な状態であれば、遷移数を +1
			if nx.GetMask(next) == 0 {
				A[next][v]++
			}
		}
	}

	// 行列 A の N 乗を計算
	ansMat := mod.powModMatrix(A, N)

	// 初期状態はルートノード (インデックス 0) にいる通りの数が 1、他は 0。
	// したがって、長さ N の文字列を作ったあとに、各安全な状態 i にいる通りの数は ansMat[i][0] となる。
	ans := 0
	for i := 0; i < M; i++ {
		if nx.GetMask(i) == 0 {
			ans = mod.add(ans, ansMat[i][0])
		}
	}

	out(ans)

}
