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

func saIs(s []int, upper int) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		if s[0] < s[1] {
			return []int{0, 1}
		} else {
			return []int{1, 0}
		}
	}
	sa := make([]int, n)
	ls := make([]bool, n)
	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			ls[i] = ls[i+1]
		} else {
			ls[i] = s[i] < s[i+1]
		}
	}
	sumL, sumS := make([]int, upper+1), make([]int, upper+1)
	for i := 0; i < n; i++ {
		if !ls[i] {
			sumS[s[i]]++
		} else {
			sumL[s[i]+1]++
		}
	}
	for i := 0; i <= upper; i++ {
		sumS[i] += sumL[i]
		if i < upper {
			sumL[i+1] += sumS[i]
		}
	}
	induce := func(lms []int) {
		for i := 0; i < n; i++ {
			sa[i] = -1
		}
		buf := make([]int, upper+1)
		copy(buf, sumS)
		for _, d := range lms {
			if d == n {
				continue
			}
			sa[buf[s[d]]] = d
			buf[s[d]]++
		}
		copy(buf, sumL)
		sa[buf[s[n-1]]] = n - 1
		buf[s[n-1]]++
		for i := 0; i < n; i++ {
			v := sa[i]
			if v >= 1 && !ls[v-1] {
				sa[buf[s[v-1]]] = v - 1
				buf[s[v-1]]++
			}
		}
		copy(buf, sumL)
		for i := n - 1; i >= 0; i-- {
			v := sa[i]
			if v >= 1 && ls[v-1] {
				buf[s[v-1]+1]--
				sa[buf[s[v-1]+1]] = v - 1
			}
		}
	}
	lmsMap := make([]int, n+1)
	for i, _ := range lmsMap {
		lmsMap[i] = -1
	}
	m := 0
	for i := 1; i < n; i++ {
		if !ls[i-1] && ls[i] {
			lmsMap[i] = m
			m++
		}
	}
	lms := make([]int, 0, m)
	for i := 1; i < n; i++ {
		if !ls[i-1] && ls[i] {
			lms = append(lms, i)
		}
	}
	induce(lms)
	if m != 0 {
		sortedLms := make([]int, 0, m)
		for _, v := range sa {
			if lmsMap[v] != -1 {
				sortedLms = append(sortedLms, v)
			}
		}
		recS := make([]int, m)
		recUpper := 0
		recS[lmsMap[sortedLms[0]]] = 0
		for i := 1; i < m; i++ {
			l := sortedLms[i-1]
			r := sortedLms[i]
			endL, endR := n, n
			if lmsMap[l]+1 < m {
				endL = lms[lmsMap[l]+1]
			}
			if lmsMap[r]+1 < m {
				endR = lms[lmsMap[r]+1]
			}
			same := true
			if endL-l != endR-r {
				same = false
			} else {
				for l < endL {
					if s[l] != s[r] {
						break
					}
					l++
					r++
				}
				if l == n || s[l] != s[r] {
					same = false
				}
			}
			if !same {
				recUpper++
			}
			recS[lmsMap[sortedLms[i]]] = recUpper
		}
		recSa := saIs(recS, recUpper)
		for i := 0; i < m; i++ {
			sortedLms[i] = lms[recSa[i]]
		}
		induce(sortedLms)
	}
	return sa
}
func suffixArray(s []int, upper int) []int {
	sa := saIs(s, upper)
	return sa
}
func suffixArrayT(s []int) []int {
	n := len(s)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(l, r int) bool { return s[l] < s[r] })
	s2 := make([]int, n)
	now := 0
	for i := 0; i < n; i++ {
		if i != 0 && s[idx[i-1]] != s[idx[i]] {
			now++
		}
		s2[idx[i]] = now
	}
	return saIs(s2, now)
}
func suffixArrayS(s string) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return saIs(s2, 255)
}
func lcpArray(s, sa []int) []int {
	n := len(s)
	rnk := make([]int, n)
	for i := 0; i < n; i++ {
		rnk[sa[i]] = i
	}
	lcp := make([]int, n-1)
	h := 0
	for i := 0; i < n; i++ {
		if h > 0 {
			h--
		}
		if rnk[i] == 0 {
			continue
		}
		j := sa[rnk[i]-1]
		for ; j+h < n && i+h < n; h++ {
			if s[j+h] != s[i+h] {
				break
			}
		}
		lcp[rnk[i]-1] = h
	}
	return lcp
}
func lcpArrayS(s string, sa []int) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return lcpArray(s2, sa)
}
func zAlgorithm(s []int) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	z := make([]int, n)
	z[0] = 0
	for i, j := 1, 0; i < n; i++ {
		if j+z[j] <= i {
			z[i] = 0
		} else {
			z[i] = zmin(j+z[j]-i, z[i-j])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if j+z[j] < i+z[i] {
			j = i
		}
	}
	z[0] = n
	return z
}
func zmin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func zAlgorithmS(s string) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return zAlgorithm(s2)
}

// -----------------------------
// ジェネリクス対応の SegTree
// -----------------------------
type Segtree[T any] struct {
	n      int
	size   int
	log    int
	d      []T
	e      func() T
	merger func(a, b T) T
}

func newSegtree[T any](v []T, e func() T, m func(a, b T) T) *Segtree[T] {
	seg := new(Segtree[T])
	seg.n = len(v)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]T, 2*seg.size)
	seg.e = e
	seg.merger = m
	for i := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		seg.d[seg.size+i] = v[i]
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.Update(i)
	}
	return seg
}

func (seg *Segtree[T]) Update(k int) {
	seg.d[k] = seg.merger(seg.d[2*k], seg.d[2*k+1])
}

func (seg *Segtree[T]) Set(p int, x T) {
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}

func (seg *Segtree[T]) Get(p int) T {
	return seg.d[p+seg.size]
}

func (seg *Segtree[T]) Prod(l, r int) T {
	sml, smr := seg.e(), seg.e()
	l += seg.size
	r += seg.size
	for l < r {
		if (l & 1) == 1 {
			sml = seg.merger(sml, seg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = seg.merger(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return seg.merger(sml, smr)
}

func (seg *Segtree[T]) AllProd() T {
	return seg.d[1]
}

func (seg *Segtree[T]) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

// type Node struct {
// 	maxVal int
// 	index  int
// }
// func main() {
// 	// (入出力の準備やNの読み込みなどは省略)
// 	a := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3} // 例
// 	N := len(a)

// 	const inf = int(1e18)

// 	e := func() Node {
// 		return Node{maxVal: -inf, index: -1}
// 	}

// 	merger := func(a, b Node) Node {
// 		if a.maxVal >= b.maxVal {
// 			return a
// 		}
// 		return b
// 	}

// 	v := make([]Node, N)
// 	for i := 0; i < N; i++ {
// 		v[i] = Node{maxVal: a[i], index: i}
// 	}

// 	seg := newSegtree[Node](v, e, merger)

// 	// 例えば区間 [2, 7) の最大値とそのインデックスを取得
// 	res := seg.Prod(2, 7)

// 	// res は Node 型で返ってくるので、そのまま中身にアクセス可能！
// 	fmt.Printf("区間内の最大値は %d (元の配列のインデックス: %d) です\n", res.maxVal, res.index)
//     // 出力結果: 区間内の最大値は 9 (元の配列のインデックス: 5) です
// }

// 宣言はこんな感じ
// seg := newSegtree[S](v, e, merger)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	Q := getI()
	s := getS()
	t := getS()

	z := t + "|" + s

	d := zAlgorithmS(z)[len(t)+1:]

	e := func() int { return 0 }
	merger := func(a, b int) int {
		return max(a, b)
	}
	seg := newSegtree(d, e, merger)

	for qi := 0; qi < Q; qi++ {
		l, r := getI()-1, getI()-len(t)+1

		c := seg.Prod(l, r)
		if c == len(t) {
			out("Yes")
		} else {
			out("No")
		}
		// out(d, l, r, c)
	}

}
