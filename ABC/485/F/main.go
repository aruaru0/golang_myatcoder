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

// Deque は動的リングバッファを用いたジェネリックな両端キュー（デック）。
// 内部配列のサイズを常に2の累乗に保つことで、ビット演算(&)による高速なインデックス折り返しを行う。
//
// 基本的な使い方:
//
//	dq := NewDeque[int](0) // 初期容量を指定して生成（0以下でも最小容量2で初期化される）
//	dq.PushBack(10)        // 末尾に追加: [10]
//	dq.PushFront(20)       // 先頭に追加: [20, 10]
//	v := dq.At(0)          // 先頭から0番目の要素を参照: 20
//	*dq.AtPtr(1) = 30      // 先頭から1番目の要素を直接書き換え: [20, 30]
//	f := dq.PopFront()     // 先頭から取り出し: 20（残り: [30]）
//	b := dq.PopBack()      // 末尾から取り出し: 30（残り: []）
type Deque[T any] struct {
	buf        []T
	head, tail int // head: 先頭要素の位置, tail: 次に末尾追加される位置
	count      int // 現在の要素数
}

// NewDeque は指定した初期容量 cap 以上となる最小の2の累乗サイズの Deque を生成する。
// おおよその最大要素数が事前に分かっている場合は cap に指定すると再確保（grow）を防げる。
func NewDeque[T any](cap int) *Deque[T] {
	size := 2
	for size < cap {
		size *= 2
	}
	return &Deque[T]{buf: make([]T, size)}
}

// Len は現在格納されている要素数を返す。O(1)
func (d *Deque[T]) Len() int { return d.count }

// grow は内部バッファが満杯になった際に容量を2倍に拡張する（内部用）。
func (d *Deque[T]) grow() {
	newBuf := make([]T, len(d.buf)*2)
	for i := 0; i < d.count; i++ {
		newBuf[i] = d.buf[(d.head+i)&(len(d.buf)-1)]
	}
	d.buf = newBuf
	d.head = 0
	d.tail = d.count
}

// PushBack は末尾に要素 v を追加する。償却 O(1)
func (d *Deque[T]) PushBack(v T) {
	if d.count == len(d.buf) {
		d.grow()
	}
	d.buf[d.tail] = v
	d.tail = (d.tail + 1) & (len(d.buf) - 1)
	d.count++
}

// PushFront は先頭に要素 v を追加する。償却 O(1)
func (d *Deque[T]) PushFront(v T) {
	if d.count == len(d.buf) {
		d.grow()
	}
	d.head = (d.head - 1) & (len(d.buf) - 1)
	d.buf[d.head] = v
	d.count++
}

// PopFront は先頭の要素を削除し、その値を返す。O(1)
// 注意: Len() == 0 の状態で呼び出すと不整合を起こすため、事前に空でないことを確認すること。
func (d *Deque[T]) PopFront() T {
	v := d.buf[d.head]
	d.head = (d.head + 1) & (len(d.buf) - 1)
	d.count--
	return v
}

// PopBack は末尾の要素を削除し、その値を返す。O(1)
// 注意: Len() == 0 の状態で呼び出すと不整合を起こすため、事前に空でないことを確認すること。
func (d *Deque[T]) PopBack() T {
	d.tail = (d.tail - 1) & (len(d.buf) - 1)
	v := d.buf[d.tail]
	d.count--
	return v
}

// At は先頭を 0 とした i 番目の要素の値を返す。O(1)
// 注意: 0 <= i < Len() の範囲外チェックは行わない。
func (d *Deque[T]) At(i int) T {
	return d.buf[(d.head+i)&(len(d.buf)-1)]
}

// AtPtr は先頭を 0 とした i 番目の要素へのポインタを返す。要素の直接更新や chmax 等に使用する。O(1)
// 注意: その後 PushFront / PushBack によりバッファ拡張(grow)が起きると古い配列を指すため、取得後は即座に使用すること。
func (d *Deque[T]) AtPtr(i int) *T {
	return &d.buf[(d.head+i)&(len(d.buf)-1)]
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()

	tot := 0
	vs := make([][]int, m+1)
	for i := 0; i < n; i++ {
		c, v := getI(), getI()
		vs[c] = append(vs[c], v)
		tot += v
	}

	ans := make([]int, m+1)
	const inf = int(1e18)

	dp := make([]*Deque[int], n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = NewDeque[int](m + 1)
		initVal := -inf
		if i == 0 {
			initVal = 0
		}
		for j := 0; j < m+1; j++ {
			dp[i].PushBack(initVal)
		}
	}

	for k := m; k >= 0; k-- {
		for _, v := range vs[k] {
			tot -= v
			for i := n - 1; i >= 0; i-- {
				for j := 0; j < m+1; j++ {
					chmax(dp[i+1].AtPtr(j), dp[i].At(j)+v)
				}
			}
		}

		now := -inf
		for i := 0; i < n+1; i++ {
			chmax(&now, dp[i].At(m))
		}
		ans[k] = now + tot

		for i := 0; i < n+1; i++ {
			steps := min(i, m+1)
			for j := 0; j < steps; j++ {
				dp[i].PushFront(-inf)
				dp[i].PopBack()
			}
		}
	}

	for i := 0; i < m+1; i++ {
		out(ans[i])
	}
}
