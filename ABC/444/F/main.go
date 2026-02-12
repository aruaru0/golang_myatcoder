package main

import (
	"bufio"
	"cmp"
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

// Queue ... 実装例
type Queue[T cmp.Ordered] []T

func (q Queue[T]) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q Queue[T]) Len() int {
	return len(q)
}

func (q Queue[T]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue[T]) Push(x T) {
	*q = append(*q, x)
	cur := q.Len()
	parent := cur / 2
	for cur != 1 {
		if q.Less(cur-1, parent-1) {
			q.Swap(cur-1, parent-1)
		} else {
			break
		}
		cur = parent
		parent = cur / 2
	}
}

func (q *Queue[T]) Pop() T {
	old := *q
	n := len(old)
	item := old[0]

	old[0] = old[n-1]
	old = old[:n-1]
	cur := 1
	for {
		nxt0 := cur * 2
		nxt1 := cur*2 + 1
		if nxt0 > len(old) {
			break
		}
		nxt := nxt0
		if nxt1 <= len(old) && old.Less(nxt1-1, nxt0-1) {
			nxt = nxt1
		}
		if old.Less(nxt-1, cur-1) {
			old.Swap(nxt-1, cur-1)
		} else {
			break
		}

		cur = nxt
	}

	*q = old
	return item
}

func solve() {
	n := getI()
	m := getI()
	a := getInts(n)

	h := (n+m)/2 + 1

	// 二分探索の判定関数
	f := func(x int) bool {
		// チェック1: 初期状態で x 以上の要素が少なすぎて、m回足しても h に届かない場合は不可
		// C++コードの最初のブロックに相当
		{
			cntBig := 0
			for i := 0; i < n; i++ {
				if a[i] >= x {
					cntBig++
				}
			}
			if cntBig+m < h {
				return false
			}
		}

		one := 0 // 長さが x 未満の合計長さ
		big := 0 // 長さが x 以上の個数
		mp := make(map[int]int)

		// add関数: 長さlenの要素をnum個追加する
		add := func(length, num int) {
			if length < x {
				one += length * num
			} else {
				mp[length] += num
				big += num
			}
		}

		// cut関数: 長さlengthの要素をnum個分割する
		// Goでは再帰的な定義が必要な場合、先に変数を宣言する
		cut := func(length, num int) {
			mp[length] -= num
			if mp[length] == 0 {
				delete(mp, length)
			}
			big -= num
			// 分割ロジック: len/2 と len - len/2 (つまり (len+1)/2)
			add(length/2, num)
			add(length-length/2, num)
		}

		// 各初期要素について分割シミュレーション（x以上を維持できる最大サイズまで分割）
		for i := 0; i < n; i++ {
			length := a[i]
			num := 1
			// 親の半分がまだ x 以上なら分割を続ける（倍々ゲーム）
			for (length+1)/2 >= x {
				length /= 2
				num *= 2
			}
			// 余り部分の処理
			r := a[i] - length*num
			l := num - r
			add(length, l)
			if r > 0 {
				add(length+1, r)
				// 特殊ケース: 長さが 2x-1 の場合、半分に割ると x と x-1 になる境界値
				if length+1 == x*2-1 {
					cut(length+1, r)
				}
			}
		}

		// x以上の個数が目標 h に届いていなければ失敗
		if big < h {
			return false
		}

		// x以上の個数が h を超えている場合、大きい方から分割して減らす（Greedy）
		// マップからキーを取り出して降順ソート（C++のmp.rbegin()の代わり）
		// cutすると常に x 未満になるため（初期ループの条件より）、再ソートは不要
		keys := make([]int, 0, len(mp))
		for k := range mp {
			keys = append(keys, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))

		for _, length := range keys {
			if big <= h {
				break
			}
			num := mp[length]
			// 減らす個数は、現在ある個数(num) と 超過分(big - h) の小さい方
			take := num
			if big-h < take {
				take = big - h
			}
			cut(length, take)
		}

		return one+big >= n+m
	}

	// 二分探索
	const INF = 1001001001
	ac := 1
	wa := INF
	for ac+1 < wa {
		wj := (ac + wa) / 2
		if f(wj) {
			ac = wj
		} else {
			wa = wj
		}
	}
	fmt.Fprintln(wr, ac)
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
