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

const inf = int(1e10)

var a []int
var N, K int

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, K = getInt(), getInt()
	a = getInts(N)
	seg := SegtreeInit(3000100, 0)

	ans := 0
	for i := 0; i < N; i++ {
		l := max(0, a[i]-K)
		r := min(a[i]+K, 3000000) + 1
		m := seg.Query(l, r)
		seg.UpdateAt(a[i], m+1)
		ans = max(ans, int(m)+1)
	}
	out(ans)
}

/*
  セグメント木(2020.05.24作成)
	SegtreeInitで初期化
	Setで値設定し、Updateで木作成
	Getで値を取得
	UpdateAtで個別のアイテムを更新
	Queryで区間の値を取得
	compareで比較方法変更可能
*/

// Data :
// Data型をstructすれば複数データが持てる
// compareも変更すること！
type Data int

// SegmentTree :
type SegmentTree struct {
	inf    Data
	d      []Data
	offset int
}

// SegtreeInit :　nが要素数、valが初期値
func SegtreeInit(n int, val Data) *SegmentTree {
	var ret SegmentTree
	size := 1
	for size < n {
		size *= 2
	}
	ret.d = make([]Data, size*2)
	for i := 1; i < size*2; i++ {
		ret.d[i] = val
	}
	ret.offset = size
	ret.inf = val
	return &ret
}

// Set : 要素に値をセット（※木は更新されない）
func (s *SegmentTree) Set(idx int, val Data) {
	s.d[s.offset+idx] = val
}

// Get : 要素に値を取得
func (s *SegmentTree) Get(idx int) Data {
	return s.d[s.offset+idx]
}

// Update :
func (s *SegmentTree) Update() {
	N := s.offset
	off := s.offset

	for N > 1 {
		for i := off; i < off+N; i += 2 {
			p := i / 2
			l := i
			r := i + 1
			s.d[p] = s.compare(s.d[l], s.d[r])
		}
		off /= 2
		N /= 2
	}
}

// querySub :
// a, b ... 範囲
func (s *SegmentTree) querySub(a, b, k, l, r int) Data {
	if r <= a || b <= l {
		return s.inf
	}
	if a <= l && r <= b {
		return s.d[k]
	}
	return s.compare(
		s.querySub(a, b, k*2, l, (l+r)/2),
		s.querySub(a, b, k*2+1, (l+r)/2, r))
}

// Query :
// a, b ... 範囲 a <= x < bの範囲で検索
// [a, b)となっているのに注意
func (s *SegmentTree) Query(a, b int) Data {
	return s.querySub(a, b, 1, 0, s.offset)
}

// UpdateAt :
func (s *SegmentTree) UpdateAt(n int, val Data) {
	pos := s.offset + n
	s.d[pos] = val
	for pos > 1 {
		p := pos / 2
		l := p * 2
		r := p*2 + 1
		s.d[p] = s.compare(s.d[l], s.d[r])
		pos /= 2
	}
}

// compare :
// 比較関数（ここで比較方法を設定）
// ※min,maxを入れ替えるときなどは、Initの設定注意
func (s *SegmentTree) compare(l, r Data) Data {
	// 区間の合計の場合はinitを0にして下記
	// return l + r

	// 区間のminの場合はinfに最大値以上を設定して下記
	if l > r {
		return l
	}
	return r
}
